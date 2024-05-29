import { getPublicClient } from '@wagmi/core';
import { formatGwei } from 'viem';

import { gasLimitConfig } from '$config';
import { PUBLIC_FEE_MULTIPLIER } from '$env/static/public';
import { NoCanonicalInfoFoundError } from '$libs/error';
import { type NFT, type Token, TokenType } from '$libs/token';
import { getTokenAddresses } from '$libs/token/getTokenAddresses';
import { getBaseFee } from '$libs/util/getBaseFee';
import { getLogger } from '$libs/util/logger';
import { config } from '$libs/wagmi';

const log = getLogger('libs:recommendedProcessingFee');

type RecommendProcessingFeeArgs = {
  token: Token | NFT;
  destChainId: number;
  srcChainId?: number;
};

export async function recommendProcessingFee({
  token,
  destChainId,
  srcChainId,
}: RecommendProcessingFeeArgs): Promise<bigint> {
  if (!srcChainId) {
    return 0n;
  }

  let estimatedMsgGaslimit;

  const baseFee = await getBaseFee(BigInt(destChainId));

  const destPublicClient = getPublicClient(config, { chainId: destChainId });

  if (!destPublicClient) throw new Error('Could not get public client');

  const maxPriorityFee = await destPublicClient.estimateMaxPriorityFeePerGas();
  log(`maxPriorityFee: ${formatGwei(maxPriorityFee)} gwei`);

  const gasPrice = await destPublicClient.getGasPrice();
  log(`gasPrice: ${formatGwei(gasPrice)} gwei`);

  if (!baseFee) throw new Error('Unable to get base fee');
  log(`baseFee: ${formatGwei(baseFee)} gwei`);

  if (token.type !== TokenType.ETH) {
    const tokenInfo = await getTokenAddresses({ token, srcChainId, destChainId });
    if (!tokenInfo) throw new NoCanonicalInfoFoundError();

    let isTokenAlreadyDeployed = false;

    if (tokenInfo.bridged) {
      const { address } = tokenInfo.bridged;
      if (address) {
        isTokenAlreadyDeployed = true;
      }
    }
    if (token.type === TokenType.ERC20) {
      if (isTokenAlreadyDeployed) {
        log(`token ${token.symbol} is already deployed on chain ${destChainId}`);

        estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE) + gasLimitConfig.erc20DeployedGasLimit;
      } else {
        log(`token ${token.symbol} is not deployed on chain ${destChainId}`);
        estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE) + gasLimitConfig.erc20NotDeployedGasLimit;
      }
    } else if (token.type === TokenType.ERC721) {
      if (isTokenAlreadyDeployed) {
        log(`token ${token.symbol} is already deployed on chain ${destChainId}`);
        estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE) + gasLimitConfig.erc721DeployedGasLimit;
      } else {
        log(`token ${token.symbol} is not deployed on chain ${destChainId}`);
        estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE) + gasLimitConfig.erc721NotDeployedGasLimit;
      }
    } else if (token.type === TokenType.ERC1155) {
      if (isTokenAlreadyDeployed) {
        log(`token ${token.symbol} is already deployed on chain ${destChainId}`);
        estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE) + gasLimitConfig.erc1155DeployedGasLimit;
      } else {
        log(`token ${token.symbol} is not deployed on chain ${destChainId}`);
        estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE) + gasLimitConfig.erc1155NotDeployedGasLimit;
      }
    }
  } else {
    log(`Fee for ETH bridging`);
    estimatedMsgGaslimit = BigInt(gasLimitConfig.GAS_RESERVE);
  }
  if (!estimatedMsgGaslimit) throw new Error('Unable to calculate fee');

  // Initial fee multiplicator and add fallback
  let feeMultiplicator: number = parseInt(PUBLIC_FEE_MULTIPLIER);

  if (gasPrice <= 50000000n) {
    feeMultiplicator = 4;
    log(`gasPrice  ${formatGwei(gasPrice)} is less than 0.5 gwei, setting feeMultiplicator to 4`);
  } else if (gasPrice <= 100000000n && gasPrice > 50000000n) {
    feeMultiplicator = 3;
    log(
      `gasPrice ${formatGwei(gasPrice)} is less than 0.1 gwei and more than 0.05 gwei, setting feeMultiplicator to 3`,
    );
  } else {
    feeMultiplicator = 2;
    log(`gasPrice ${formatGwei(gasPrice)} is more than 0.1 gwei, setting feeMultiplicator to 2`);
  }

  const fee = estimatedMsgGaslimit * gasPrice * BigInt(feeMultiplicator);
  return fee;
}

// function roundWeiTo6DecimalPlaces(wei: bigint): bigint {
//   const roundingFactor = BigInt('1000000000000'); // 10^12

//   // Calculate how many "10^12 wei" units are in the input
//   const units = wei / roundingFactor;

//   // Multiply back to get the rounded wei value
//   const roundedWei = units * roundingFactor;
//   return roundedWei;
// }
