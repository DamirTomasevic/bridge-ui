import { writable } from 'svelte/store';
import type { Signer, Transaction, ethers } from 'ethers';
import type { BridgeTransaction } from '../domain/transactions';
import { Deferred } from '../utils/Deferred';

export const transactions = writable<BridgeTransaction[]>([]);

// Custom store: pendingTransactions
const { subscribe, set, update } = writable<Transaction[]>([]);
export const pendingTransactions = {
  /**
   * We're creating here a custom store, which is a writable store.
   * We must stick to the store contract, which is:
   */
  set,
  subscribe,
  // update, // this method is optional.

  /**
   * Custom method, which will help us add a new transaction to the store
   * and get it removed onces the transaction is mined.
   */
  add: (tx: Transaction, signer: Signer) => {
    const deferred = new Deferred<ethers.providers.TransactionReceipt>();

    update((txs: Transaction[]) => {
      // New array with the new transaction appended
      const newPendingTransactions = [...txs, tx];

      // Save the index of the new transaction to later on remove it
      // from the list of pending transactions.
      const idxAppendedTransaction = newPendingTransactions.length - 1;

      // Next step is to wait for the transaction to be mined
      // before removing it from the store.

      /**
       * Returns a Promise which will not resolve until transactionHash is mined.
       * If confirms is 0, this method is non-blocking and if the transaction
       * has not been mined returns null. Otherwise, this method will block until
       * the transaction has confirms blocks mined on top of the block in which
       * is was mined.
       * See https://docs.ethers.org/v5/api/providers/provider/#Provider-waitForTransaction
       */
      signer.provider.waitForTransaction(tx.hash, 1).then((receipt) => {
        // The transaction has been mined.

        // Removes the transaction from the store
        update((txs: Transaction[]) => {
          const copyPendingTransactions = [...txs];
          copyPendingTransactions.splice(idxAppendedTransaction, 1);
          return copyPendingTransactions;
        });

        // Resolves or rejects the promise depending on the transaction status.
        if (receipt.status === 1) {
          deferred.resolve(receipt);
        } else {
          deferred.reject(receipt);
        }
      });

      return newPendingTransactions;
    });

    return deferred.promise;
  },
};
