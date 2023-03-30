import { BigNumber, Contract, ethers } from 'ethers';
import type { Transaction } from 'ethers';
import type {
  ApproveOpts,
  Bridge,
  BridgeOpts,
  ClaimOpts,
  ReleaseOpts,
} from '../domain/bridge';
import type { Prover } from '../domain/proof';
import BridgeABI from '../constants/abi/Bridge';
import { chains } from '../chain/chains';
import { type Message, MessageStatus } from '../domain/message';

export class ETHBridge implements Bridge {
  private readonly prover: Prover;

  constructor(prover: Prover) {
    this.prover = prover;
  }

  static async prepareTransaction(
    opts: BridgeOpts,
  ): Promise<{ contract: Contract; message: Message; owner: string }> {
    const contract: Contract = new Contract(
      opts.bridgeAddress,
      BridgeABI,
      opts.signer,
    );

    const owner = await opts.signer.getAddress();
    const message: Message = {
      sender: owner,
      srcChainId: opts.fromChainId,
      destChainId: opts.toChainId,
      owner: owner,
      to: opts.to,
      refundAddress: owner,
      depositValue:
        opts.to.toLowerCase() === owner.toLowerCase()
          ? opts.amountInWei
          : BigNumber.from(0),
      callValue:
        opts.to.toLowerCase() === owner.toLowerCase()
          ? BigNumber.from(0)
          : opts.amountInWei,
      processingFee: opts.processingFeeInWei ?? BigNumber.from(0),
      gasLimit: opts.processingFeeInWei
        ? BigNumber.from(140000)
        : BigNumber.from(0),
      memo: opts.memo ?? '',
      id: 1, // will be set in contract,
      data: '0x',
    };

    return { contract, owner, message };
  }

  RequiresAllowance(opts: ApproveOpts): Promise<boolean> {
    return Promise.resolve(false);
  }

  // ETH does not need to be approved for transacting
  Approve(opts: ApproveOpts): Promise<Transaction> {
    return new Promise((resolve) => resolve({} as unknown as Transaction));
  }

  async Bridge(opts: BridgeOpts): Promise<Transaction> {
    const { contract, message } = await ETHBridge.prepareTransaction(opts);

    const tx = await contract.sendMessage(message, {
      value: message.depositValue
        .add(message.processingFee)
        .add(message.callValue),
    });

    return tx;
  }

  async EstimateGas(opts: BridgeOpts): Promise<BigNumber> {
    const { contract, message } = await ETHBridge.prepareTransaction(opts);

    const gasEstimate = await contract.estimateGas.sendMessage(message, {
      value: message.depositValue
        .add(message.processingFee)
        .add(message.callValue),
    });

    return gasEstimate;
  }

  async Claim(opts: ClaimOpts): Promise<Transaction> {
    const contract: Contract = new Contract(
      opts.destBridgeAddress,
      BridgeABI,
      opts.signer,
    );

    const messageStatus: MessageStatus = await contract.getMessageStatus(
      opts.msgHash,
    );

    if (messageStatus === MessageStatus.Done) {
      throw Error('message already processed');
    }

    const signerAddress = await opts.signer.getAddress();

    if (opts.message.owner.toLowerCase() !== signerAddress.toLowerCase()) {
      throw Error('user can not process this, it is not their message');
    }

    if (messageStatus === MessageStatus.New) {
      const proofOpts = {
        srcChain: opts.message.srcChainId,
        msgHash: opts.msgHash,
        sender: opts.srcBridgeAddress,
        srcBridgeAddress: opts.srcBridgeAddress,
        destChain: opts.message.destChainId,
        destHeaderSyncAddress:
          chains[opts.message.destChainId].headerSyncAddress,
        srcSignalServiceAddress:
          chains[opts.message.srcChainId].signalServiceAddress,
      };

      const proof = await this.prover.generateProof(proofOpts);
      let processMessageTx;
      try {
        processMessageTx = await contract.processMessage(opts.message, proof);
      } catch (error) {
        if (error.code === ethers.errors.UNPREDICTABLE_GAS_LIMIT) {
          processMessageTx = await contract.processMessage(
            opts.message,
            proof,
            {
              gasLimit: 1e6,
            },
          );
        } else {
          throw new Error(error);
        }
      }
      return processMessageTx;
    } else {
      return await contract.retryMessage(opts.message, true);
    }
  }

  async ReleaseTokens(opts: ReleaseOpts): Promise<Transaction> {
    const destBridgeContract: Contract = new Contract(
      opts.destBridgeAddress,
      BridgeABI,
      opts.destProvider,
    );

    const messageStatus: MessageStatus =
      await destBridgeContract.getMessageStatus(opts.msgHash);

    if (messageStatus === MessageStatus.Done) {
      throw Error('message already processed');
    }

    const signerAddress = await opts.signer.getAddress();

    if (opts.message.owner.toLowerCase() !== signerAddress.toLowerCase()) {
      throw Error('user can not release these tokens, it is not their message');
    }

    if (messageStatus === MessageStatus.Failed) {
      const proofOpts = {
        srcChain: opts.message.srcChainId,
        msgHash: opts.msgHash,
        sender: opts.srcBridgeAddress,
        destBridgeAddress: opts.destBridgeAddress,
        destChain: opts.message.destChainId,
        destHeaderSyncAddress:
          chains[opts.message.destChainId].headerSyncAddress,
        srcHeaderSyncAddress: chains[opts.message.srcChainId].headerSyncAddress,
      };

      const proof = await this.prover.generateReleaseProof(proofOpts);

      const srcBridgeContract: Contract = new Contract(
        opts.srcBridgeAddress,
        BridgeABI,
        opts.signer,
      );

      return await srcBridgeContract.releaseEther(opts.message, proof);
    }
  }
}
