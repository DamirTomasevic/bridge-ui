import { readContract } from '@wagmi/core';

import { taikoonTokenAbi, taikoonTokenAddress } from '../../generated/abi';
import getConfig from '../wagmi/getConfig';

export async function maxSupply(): Promise<number> {
  const { config, chainId } = getConfig();

  const result = await readContract(config, {
    abi: taikoonTokenAbi,
    address: taikoonTokenAddress[chainId],
    functionName: 'maxSupply',
    chainId,
  });

  return parseInt(result.toString(16), 16);
}
