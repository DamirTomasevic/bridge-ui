import { fetchGuardianProversFromContract } from './guardianProver/fetchGuardianProversFromContract';
import { GuardianProverStatus, type Guardian, type SignedBlocks } from './types';
import { fetchSignedBlocksFromApi } from './api/signedBlocksApiCalls';
import { getGuardianProverIdsPerBlockNumber } from './blocks/getGuardianProverIdsPerBlockNumber';
import { sortSignedBlocksDescending } from './blocks/sortSignedBlocks';
import { publicClient } from './wagmi/publicClient';
import { formatEther, type Address } from 'viem';
import { fetchLatestGuardianProverHealtCheckFromApi, fetchUptimeFromApi } from './api';
import { fetchGuardianProverRequirementsFromContract } from './guardianProver/fetchGuardianProverRequirementsFromContract';
import {
	minGuardianRequirement,
	lastGuardianFetchTimestamp,
	guardianProvers,
	signedBlocks,
	signerPerBlock,
	loading,
	totalGuardianProvers
} from '$stores';
import { get, writable } from 'svelte/store';

const BLOCKS_TO_CHECK = 20;
const THRESHOLD = BLOCKS_TO_CHECK / 2;
const HEALTHCHECK_TIMEOUT_IN_SECONDS = 60;

const tempGuardianStore = writable<Guardian[]>([]);

let guardiansIntervalId;
let blocksAndLivelinessIntervalId;

export async function startFetching() {
	if (get(loading)) return;

	await refreshData();

	guardiansIntervalId = setInterval(() => {
		fetchGuardians();
	}, 10000);

	blocksAndLivelinessIntervalId = setInterval(() => {
		fetchSignedBlockStats();
		determineLiveliness();
	}, 12000);
}

export function stopFetching() {
	if (guardiansIntervalId) {
		clearInterval(guardiansIntervalId);
		guardiansIntervalId = null;
	}
	if (blocksAndLivelinessIntervalId) {
		clearInterval(blocksAndLivelinessIntervalId);
		blocksAndLivelinessIntervalId = null;
	}
}

export async function refreshData() {
	if (get(loading) === true) return;
	loading.set(true);

	if (get(guardianProvers)?.length === 0) {
		await fetchGuardians();
		const block = fetchSignedBlockStats();
		const liveness = determineLiveliness();
		await Promise.all([block, liveness]);
	} else {
		const guardian = fetchGuardians();
		const block = fetchSignedBlockStats();
		const liveness = determineLiveliness();
		await Promise.all([block, guardian, liveness]);
	}

	loading.set(false);
}

async function fetchGuardians() {
	const [rawData, required] = await Promise.all([
		fetchGuardianProversFromContract(),
		fetchGuardianProverRequirementsFromContract()
	]);
	minGuardianRequirement.set(required);
	totalGuardianProvers.set(rawData.length);

	const guardianFetchPromises = rawData.map(async (guardian) => {
		const balance = await publicClient.getBalance({
			address: guardian.address as Address
		});

		const balanceAsEther = formatEther(balance);
		guardian.balance = balanceAsEther;

		const [status, uptime] = await Promise.all([
			fetchLatestGuardianProverHealtCheckFromApi(
				import.meta.env.VITE_GUARDIAN_PROVER_API_URL,
				guardian.id
			),
			fetchUptimeFromApi(import.meta.env.VITE_GUARDIAN_PROVER_API_URL, guardian.id)
		]);

		guardian.latestHealthCheck = status;
		guardian.uptime = Math.min(uptime, 100);
		// guardian.alive = status.alive ? GuardianProverStatus.ALIVE : GuardianProverStatus.DEAD;

		return guardian;
	});

	const data = await Promise.all(guardianFetchPromises);
	tempGuardianStore.set(data);

	lastGuardianFetchTimestamp.set(Date.now());
}

async function fetchSignedBlockStats() {
	const blocks: SignedBlocks = await fetchSignedBlocksFromApi(
		import.meta.env.VITE_GUARDIAN_PROVER_API_URL
	);

	signedBlocks.set(sortSignedBlocksDescending(blocks));

	const signer = await getGuardianProverIdsPerBlockNumber(blocks);
	signerPerBlock.set(signer);
}

async function determineLiveliness(): Promise<void> {
	const tempData = get(tempGuardianStore);
	const now = new Date();
	if (!tempData) return;

	const guardians = tempData.map((guardian) => {
		const latestCheck = guardian.latestHealthCheck;
		const createdAt = new Date(latestCheck.createdAt);
		const secondsSinceLastCheck = (now.getTime() - createdAt.getTime()) / 1000;

		if (secondsSinceLastCheck > HEALTHCHECK_TIMEOUT_IN_SECONDS) {
			return { ...guardian, alive: GuardianProverStatus.DEAD };
		}

		let countSignedBlocks = 0;
		const recentSignedBlocks = get(signedBlocks).slice(0, BLOCKS_TO_CHECK);

		for (const block of recentSignedBlocks) {
			if (block.blocks.some((b) => b.guardianProverID === Number(guardian.id))) {
				countSignedBlocks++;
			}
		}

		const status =
			countSignedBlocks >= THRESHOLD ? GuardianProverStatus.ALIVE : GuardianProverStatus.UNHEALTHY;

		return {
			...guardian,
			alive: status
		};
	});

	guardianProvers.set(guardians);
}
