import { readContract } from '@wagmi/core';

import { config } from '$wagmi-config';

import { taikoonTokenAbi, taikoonTokenAddress } from '../../generated/abi/';
import { web3modal } from '../../lib/connect';
import type { IAddress, IChainId } from '../../types';

export async function balanceOf(address: IAddress): Promise<number> {
  const { selectedNetworkId } = web3modal.getState();
  if (!selectedNetworkId) return 0;

  const chainId = selectedNetworkId as IChainId;

  const result = await readContract(config, {
    abi: taikoonTokenAbi,
    address: taikoonTokenAddress[chainId],
    functionName: 'balanceOf',
    args: [address],
    chainId,
  });

  return parseInt(result.toString());
}
