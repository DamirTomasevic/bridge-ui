export const recommentProcessingFee = {
  ethGasLimit: BigInt(900_000),
  erc20NotDeployedGasLimit: BigInt(3_100_000),
  erc20DeployedGasLimit: BigInt(1_100_000),
};

export const processingFeeComponent = {
  closingDelayOptionClick: 300,
  intervalComputeRecommendedFee: 20_000,
};

export const bridgeService = {
  noOwnerGasLimit: BigInt(140_000),
  noTokenDeployedGasLimit: BigInt(3_000_000),
  erc20GasLimitThreshold: BigInt(2_500_000),
  unpredictableGasLimit: BigInt(1_000_000),
};

export const pendingTransaction = {
  waitTimeout: 300_000,
};

export const storageService = {
  bridgeTxPrefix: 'transactions',
  customTokenPrefix: 'custom-tokens',
};

export const bridgeTransactionPoller = {
  interval: 20_000,
};

export const statusComponent = {
  minimumEthToClaim: 0.0001,
};

export const transactionConfig = {
  pageSizeDesktop: 6,
  pageSizeMobile: 4,
  blurTransitionTime: 300,
};
