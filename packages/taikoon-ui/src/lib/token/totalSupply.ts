import { readContract } from '@wagmi/core';

import { taikoonTokenAbi, taikoonTokenAddress } from '../../generated/abi';
import getConfig from '../../lib/wagmi/getConfig';

export async function totalSupply(): Promise<number> {
  const { config, chainId } = getConfig();

  const result = await readContract(config, {
    abi: taikoonTokenAbi,
    address: taikoonTokenAddress[chainId],
    functionName: 'totalSupply',
    chainId,
  });

  return parseInt(result.toString(16), 16);
}
