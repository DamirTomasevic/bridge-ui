import { watchAccount } from '@wagmi/core';

import { config } from '$wagmi-config';

import { isSupportedChain } from '../../lib/chain';
import { refreshUserBalance } from '../../lib/util/balance';
import { account } from '../../stores/account';
import { switchChainModal } from '../../stores/modal';
import { connectedSourceChain } from '../../stores/network';

let isWatching = false;
let unWatchAccount: () => void;

export async function startWatching() {
  if (!isWatching) {
    unWatchAccount = watchAccount(config, {
      onChange(data) {
        const { chain } = data;
        account.set(data);
        refreshUserBalance();
        // We need to check if the chain is supported, and if not
        // we present the user with a modal to switch networks.
        if ((chain && !isSupportedChain(Number(chain.id))) || (!data.chainId && data.address)) {
          switchChainModal.set(true);
          return;
        } else if (chain) {
          // When we switch networks, we are actually selecting
          // the source chain.
          connectedSourceChain.set(chain);
        }
      },
    });

    isWatching = true;
  }
}

export function stopWatching() {
  unWatchAccount && unWatchAccount();
  isWatching = false;
}
