import { Contract, ethers } from 'ethers';
import { RLP } from 'ethers/lib/utils.js';
import HeaderSyncABI from '../constants/abi/ICrossChainSync';
import type { Block, BlockHeader } from '../domain/block';
import type {
  Prover,
  GenerateProofOpts,
  EthGetProofResponse,
  GenerateReleaseProofOpts,
} from '../domain/proof';

export class ProofService implements Prover {
  private readonly providers: Record<
    number,
    ethers.providers.StaticJsonRpcProvider
  >;

  constructor(
    providers: Record<number, ethers.providers.StaticJsonRpcProvider>,
  ) {
    this.providers = providers;
  }

  private static getKey(opts: GenerateProofOpts | GenerateReleaseProofOpts) {
    const key = ethers.utils.keccak256(
      ethers.utils.solidityPack(
        ['address', 'bytes32'],
        [opts.sender, opts.msgHash],
      ),
    );

    return key;
  }

  private static async getBlockAndBlockHeader(
    contract: ethers.Contract,
    provider: ethers.providers.StaticJsonRpcProvider,
  ): Promise<{ block: Block; blockHeader: BlockHeader }> {
    const latestSyncedHeader = await contract.getCrossChainBlockHash(0);

    const block: Block = await provider.send('eth_getBlockByHash', [
      latestSyncedHeader,
      false,
    ]);

    const logsBloom = block.logsBloom.toString().substring(2);

    const blockHeader: BlockHeader = {
      parentHash: block.parentHash,
      ommersHash: block.sha3Uncles,
      beneficiary: block.miner,
      stateRoot: block.stateRoot,
      transactionsRoot: block.transactionsRoot,
      receiptsRoot: block.receiptsRoot,
      logsBloom: logsBloom.match(/.{1,64}/g)!.map((s: string) => '0x' + s),
      difficulty: block.difficulty,
      height: block.number,
      gasLimit: block.gasLimit,
      gasUsed: block.gasUsed,
      timestamp: block.timestamp,
      extraData: block.extraData,
      mixHash: block.mixHash,
      nonce: block.nonce,
      baseFeePerGas: block.baseFeePerGas ? parseInt(block.baseFeePerGas) : 0,
      withdrawalsRoot: block.withdrawalsRoot ?? ethers.constants.HashZero,
    };

    return { block, blockHeader };
  }

  private static getSignalProof(
    proof: EthGetProofResponse,
    blockHeader: BlockHeader,
  ) {
    // RLP encode the proof together for LibTrieProof to decode
    const encodedProof = RLP.encode(proof.storageProof[0].proof);

    // encode the SignalProof struct from LibBridgeSignal
    const signalProof = ethers.utils.defaultAbiCoder.encode(
      [
        'tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas, bytes32 withdrawalsRoot) header, bytes proof)',
      ],
      [{ header: blockHeader, proof: encodedProof }],
    );

    return signalProof;
  }

  async generateProof(opts: GenerateProofOpts): Promise<string> {
    const key = ProofService.getKey(opts);

    const provider = this.providers[opts.srcChain];

    const contract = new Contract(
      opts.destCrossChainSyncAddress,
      HeaderSyncABI,
      this.providers[opts.destChain],
    );

    const { block, blockHeader } = await ProofService.getBlockAndBlockHeader(
      contract,
      provider,
    );

    // rpc call to get the merkle proof what value is at key on the SignalService contract
    const proof: EthGetProofResponse = await provider.send('eth_getProof', [
      opts.srcSignalServiceAddress,
      [key],
      block.hash,
    ]);

    if (proof.storageProof[0].value !== '0x1') {
      throw Error('invalid proof');
    }

    const p = ProofService.getSignalProof(proof, blockHeader);
    return p;
  }

  async generateReleaseProof(opts: GenerateReleaseProofOpts): Promise<string> {
    const key = ProofService.getKey(opts);

    const provider = this.providers[opts.destChain];

    const contract = new Contract(
      opts.srcCrossChainSyncAddress,
      HeaderSyncABI,
      this.providers[opts.srcChain],
    );

    const { block, blockHeader } = await ProofService.getBlockAndBlockHeader(
      contract,
      provider,
    );

    // rpc call to get the merkle proof what value is at key on the SignalService contract
    const proof: EthGetProofResponse = await provider.send('eth_getProof', [
      opts.destBridgeAddress,
      [key],
      block.hash,
    ]);

    if (proof.storageProof[0].value !== '0x3') {
      throw Error('invalid proof');
    }

    const p = ProofService.getSignalProof(proof, blockHeader);
    return p;
  }
}
