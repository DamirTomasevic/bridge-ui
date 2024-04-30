import { readContract } from '@wagmi/core';
import { zeroAddress } from 'viem';

import { config } from '$wagmi-config';

import { taikoonTokenAbi, taikoonTokenAddress } from '../../generated/abi/';
import { web3modal } from '../../lib/connect';
import type { IAddress, IChainId } from '../../types';

export async function ownerOf(tokenId: number): Promise<IAddress> {
  const { selectedNetworkId } = web3modal.getState();
  if (!selectedNetworkId) return zeroAddress;

  const chainId = selectedNetworkId as IChainId;

  const result = await readContract(config, {
    abi: taikoonTokenAbi,
    address: taikoonTokenAddress[chainId],
    functionName: 'ownerOf',
    args: [BigInt(tokenId)],
    chainId,
  });

  return result as IAddress;
}
