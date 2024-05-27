import { whitelist } from '$lib/whitelist';

import getConfig from '../../lib/wagmi/getConfig';

export async function maxSupply(): Promise<number> {
  const { chainId } = getConfig();

  const currentWhitelist = whitelist[chainId];

  let totalCount = 0;

  currentWhitelist.values.forEach((item: any) => {
    totalCount += parseInt(item.value[1]);
  });

  return totalCount;
}
