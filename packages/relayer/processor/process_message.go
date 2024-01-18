package processor

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/taikoxyz/taiko-mono/packages/relayer"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/bridge"
	"github.com/taikoxyz/taiko-mono/packages/relayer/bindings/icrosschainsync"
	"github.com/taikoxyz/taiko-mono/packages/relayer/pkg/proof"
	"github.com/taikoxyz/taiko-mono/packages/relayer/pkg/queue"
)

var (
	errUnprocessable = errors.New("message is unprocessable")
)

func (p *Processor) eventStatusFromMsgHash(
	ctx context.Context,
	gasLimit *big.Int,
	signal [32]byte,
) (relayer.EventStatus, error) {
	var eventStatus relayer.EventStatus

	ctx, cancel := context.WithTimeout(ctx, p.ethClientTimeout)

	defer cancel()

	messageStatus, err := p.destBridge.MessageStatus(&bind.CallOpts{
		Context: ctx,
	}, signal)
	if err != nil {
		return 0, errors.Wrap(err, "svc.destBridge.MessageStatus")
	}

	eventStatus = relayer.EventStatus(messageStatus)
	if eventStatus == relayer.EventStatusNew {
		if gasLimit == nil || gasLimit.Cmp(common.Big0) == 0 {
			// if gasLimit is 0, relayer can not process this.
			eventStatus = relayer.EventStatusNewOnlyOwner
		}
	}

	return eventStatus, nil
}

// processMessage prepares and calls `processMessage` on the bridge.
// the proof must be generated from the gethclient's eth_getProof via the Prover,
// then rlp-encoded and combined as a singular byte slice,
// then abi encoded into a SignalProof struct as the contract
// expects
func (p *Processor) processMessage(
	ctx context.Context,
	msg queue.Message,
) error {
	msgBody := &queue.QueueMessageBody{}
	if err := json.Unmarshal(msg.Body, msgBody); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	if msgBody.Event.Message.GasLimit == nil || msgBody.Event.Message.GasLimit.Cmp(common.Big0) == 0 {
		return errors.New("only user can process this, gasLimit set to 0")
	}

	eventStatus, err := p.eventStatusFromMsgHash(ctx, msgBody.Event.Message.GasLimit, msgBody.Event.MsgHash)
	if err != nil {
		return errors.Wrap(err, "p.eventStatusFromMsgHash")
	}

	if !canProcessMessage(ctx, eventStatus, msgBody.Event.Message.Owner, p.relayerAddr) {
		return errUnprocessable
	}

	if err := p.waitForConfirmations(ctx, msgBody.Event.Raw.TxHash, msgBody.Event.Raw.BlockNumber); err != nil {
		return errors.Wrap(err, "p.waitForConfirmations")
	}

	var blockNum uint64 = msgBody.Event.Raw.BlockNumber

	// wait for srcChain => destChain header to sync if no hops,
	// or srcChain => hopChain => hopChain => hopChain => destChain if hops exist.
	if p.hops != nil {
		var hopEthClient ethClient = p.srcEthClient

		for _, hop := range p.hops {
			hop.blockNum = blockNum

			_, err := p.waitHeaderSynced(ctx, hop.headerSyncer, hopEthClient, blockNum)

			if err != nil {
				return errors.Wrap(err, "p.waitHeaderSynced")
			}

			// todo: instead of latest, need way to find out which block num on the hop chain
			// the previous blockHash was synced in, and then wait for that header to be synced
			// on the next hop chain.
			snippet, err := hop.headerSyncer.GetSyncedSnippet(&bind.CallOpts{
				Context: ctx,
			},
				hop.blockNum,
			)

			slog.Info("hop synced snippet",
				"syncedInBlock", snippet.SyncedInBlock,
				"blockNum", hop.blockNum,
				"blockHash", common.Bytes2Hex(snippet.BlockHash[:]),
			)

			if err != nil {
				return errors.Wrap(err, "hop.headerSyncer.GetSyncedSnippet")
			}

			blockNum = snippet.SyncedInBlock

			hopEthClient = hop.ethClient
		}

		blockNum, err = p.waitHeaderSynced(ctx, p.destHeaderSyncer, hopEthClient, blockNum)
		if err != nil {
			return errors.Wrap(err, "p.waitHeaderSynced")
		}
	} else {
		if _, err := p.waitHeaderSynced(ctx, p.destHeaderSyncer, p.srcEthClient, msgBody.Event.Raw.BlockNumber); err != nil {
			return errors.Wrap(err, "p.waitHeaderSynced")
		}
	}

	key, err := p.srcSignalService.GetSignalSlot(&bind.CallOpts{},
		msgBody.Event.Message.SrcChainId,
		msgBody.Event.Raw.Address,
		msgBody.Event.MsgHash,
	)

	if err != nil {
		return errors.Wrap(err, "p.srcSignalService.GetSignalSlot")
	}

	hops := []proof.HopParams{}

	var encodedSignalProof []byte

	var latestSyncedSnippet icrosschainsync.ICrossChainSyncSnippet

	// if a hop is set, the proof service needs to generate an additional proof
	// for the signal service intermediary chain in between the source chain
	// and the destination chain.
	for _, hop := range p.hops {
		slog.Info(
			"adding hop",
			"hopChainId", hop.chainID.Uint64(),
			"hopSignalServiceAddress", hop.signalServiceAddress.Hex(),
		)

		hops = append(hops, proof.HopParams{
			ChainID:              hop.chainID,
			SignalServiceAddress: hop.signalServiceAddress,
			Blocker:              hop.ethClient,
			Caller:               hop.caller,
			SignalService:        hop.signalService,
			TaikoAddress:         hop.taikoAddress,
			BlockNumber:          blockNum,
		})
	}

	if len(hops) != 0 {
		encodedSignalProof, _, err = p.prover.EncodedSignalProofWithHops(
			ctx,
			p.srcCaller,
			p.srcSignalServiceAddress,
			p.destHeaderSyncAddress,
			hops,
			common.Bytes2Hex(key[:]),
			msgBody.Event.Raw.BlockHash,
			blockNum,
		)
	} else {
		// get latest synced header since not every header is synced from L1 => L2,
		// and later blocks still have the storage trie proof from previous blocks.
		latestSyncedSnippet, err = p.destHeaderSyncer.GetSyncedSnippet(&bind.CallOpts{}, 0)
		if err != nil {
			return errors.Wrap(err, "taiko.GetSyncedSnippet")
		}

		encodedSignalProof, err = p.prover.EncodedSignalProof(
			ctx,
			p.srcCaller,
			p.srcSignalServiceAddress,
			p.destHeaderSyncAddress,
			common.Bytes2Hex(key[:]),
			latestSyncedSnippet.BlockHash,
		)
	}

	if err != nil {
		slog.Error("error encoding signal proof",
			"srcChainID", msgBody.Event.Message.SrcChainId,
			"destChainID", msgBody.Event.Message.DestChainId,
			"txHash", msgBody.Event.Raw.TxHash.Hex(),
			"msgHash", common.Hash(msgBody.Event.MsgHash).Hex(),
			"from", msgBody.Event.Message.From.Hex(),
			"owner", msgBody.Event.Message.Owner.Hex(),
			"error", err,
			"hopsLength", len(hops),
		)

		return errors.Wrap(err, "p.prover.GetEncodedSignalProof")
	}

	// check if message is received first. if not, it will definitely fail,
	// so we can exit early on this one. there is most likely
	// an issue with the signal generation.
	received, err := p.destBridge.ProveMessageReceived(&bind.CallOpts{
		Context: ctx,
	}, msgBody.Event.Message, encodedSignalProof)
	if err != nil {
		return errors.Wrap(err, "p.destBridge.ProveMessageReceived")
	}

	// message will fail when we try to process it
	if !received {
		slog.Warn("Message not received on dest chain",
			"msgHash", common.Hash(msgBody.Event.MsgHash).Hex(),
			"srcChainId", msgBody.Event.Message.SrcChainId,
		)

		relayer.MessagesNotReceivedOnDestChain.Inc()

		return errors.New("message not received")
	}

	var tx *types.Transaction

	sendTx := func() error {
		if ctx.Err() != nil {
			return nil
		}

		tx, err = p.sendProcessMessageCall(ctx, msgBody.Event, encodedSignalProof)
		if err != nil {
			return err
		}

		return nil
	}

	if err := backoff.Retry(sendTx, backoff.WithMaxRetries(
		backoff.NewConstantBackOff(p.backOffRetryInterval),
		p.backOffMaxRetries),
	); err != nil {
		return err
	}

	relayer.EventsProcessed.Inc()

	ctx, cancel := context.WithTimeout(ctx, 4*time.Minute)

	defer cancel()

	receipt, err := relayer.WaitReceipt(ctx, p.destEthClient, tx.Hash())
	if err != nil {
		return errors.Wrap(err, "relayer.WaitReceipt")
	}

	if err := p.saveMessageStatusChangedEvent(ctx, receipt, msgBody.Event); err != nil {
		return errors.Wrap(err, "p.saveMEssageStatusChangedEvent")
	}

	slog.Info("Mined tx", "txHash", hex.EncodeToString(tx.Hash().Bytes()))

	messageStatus, err := p.destBridge.MessageStatus(&bind.CallOpts{}, msgBody.Event.MsgHash)
	if err != nil {
		return errors.Wrap(err, "p.destBridge.GetMessageStatus")
	}

	slog.Info(
		"updating message status",
		"status", relayer.EventStatus(messageStatus).String(),
		"occuredtxHash", msgBody.Event.Raw.TxHash.Hex(),
		"processedTxHash", hex.EncodeToString(tx.Hash().Bytes()),
	)

	if messageStatus == uint8(relayer.EventStatusRetriable) {
		relayer.RetriableEvents.Inc()
	} else if messageStatus == uint8(relayer.EventStatusDone) {
		relayer.DoneEvents.Inc()
	}

	// update message status
	if err := p.eventRepo.UpdateStatus(ctx, msgBody.ID, relayer.EventStatus(messageStatus)); err != nil {
		return errors.Wrap(err, fmt.Sprintf("p.eventRepo.UpdateStatus, id: %v", msgBody.ID))
	}

	return nil
}

func (p *Processor) sendProcessMessageCall(
	ctx context.Context,
	event *bridge.BridgeMessageSent,
	proof []byte,
) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(p.ecdsaKey, new(big.Int).SetUint64(event.Message.DestChainId))
	if err != nil {
		return nil, errors.Wrap(err, "bind.NewKeyedTransactorWithChainID")
	}

	auth.Context = ctx

	p.mu.Lock()
	defer p.mu.Unlock()

	err = p.getLatestNonce(ctx, auth)
	if err != nil {
		return nil, errors.New("p.getLatestNonce")
	}

	eventType, canonicalToken, _, err := relayer.DecodeMessageSentData(event)
	if err != nil {
		return nil, errors.Wrap(err, "relayer.DecodeMessageSentData")
	}

	var gas uint64

	var cost *big.Int

	needsContractDeployment, err := p.needsContractDeployment(ctx, event, eventType, canonicalToken)
	if err != nil {
		return nil, errors.Wrap(err, "p.needsContractDeployment")
	}

	if needsContractDeployment {
		auth.GasLimit = 3000000
	} else {
		// otherwise we can estimate gas
		gas, err = p.estimateGas(ctx, event.Message, proof)
		// and if gas estimation failed, we just try to hardcore a value no matter what type of event,
		// or whether the contract is deployed.
		if err != nil || gas == 0 {
			slog.Info("gas estimation failed, hardcoding gas limit", "p.estimateGas:", err)
			err = p.hardcodeGasLimit(ctx, auth, event, eventType, canonicalToken)
			if err != nil {
				return nil, errors.Wrap(err, "p.hardcodeGasLimit")
			}
		} else {
			auth.GasLimit = gas
		}
	}

	if err = p.setGasTipOrPrice(ctx, auth); err != nil {
		return nil, errors.Wrap(err, "p.setGasTipOrPrice")
	}

	cost, err = p.getCost(ctx, auth)
	if err != nil {
		return nil, errors.Wrap(err, "p.getCost")
	}

	if bool(p.profitableOnly) {
		profitable, err := p.isProfitable(ctx, event.Message, cost)
		if err != nil || !profitable {
			return nil, relayer.ErrUnprofitable
		}
	}

	// process the message on the destination bridge.
	tx, err := p.destBridge.ProcessMessage(auth, event.Message, proof)
	if err != nil {
		return nil, errors.Wrap(err, "p.destBridge.ProcessMessage")
	}

	p.setLatestNonce(tx.Nonce())

	return tx, nil
}

// node is unable to estimate gas correctly for contract deployments, we need to check if the token
// is deployed, and always hardcode in this case. we need to check this before calling
// estimategas, as the node will soemtimes return a gas estimate for a contract deployment, however,
// it is incorrect and the tx will revert.
func (p *Processor) needsContractDeployment(
	ctx context.Context,
	event *bridge.BridgeMessageSent,
	eventType relayer.EventType,
	canonicalToken relayer.CanonicalToken,
) (bool, error) {
	if eventType == relayer.EventTypeSendETH {
		return false, nil
	}

	var bridgedAddress common.Address

	var err error

	chainID := new(big.Int).SetUint64(canonicalToken.ChainID())
	addr := canonicalToken.Address()

	ctx, cancel := context.WithTimeout(ctx, p.ethClientTimeout)
	defer cancel()

	opts := &bind.CallOpts{
		Context: ctx,
	}

	destChainID := new(big.Int).SetUint64(event.Message.DestChainId)
	if eventType == relayer.EventTypeSendERC20 && destChainID.Cmp(chainID) != 0 {
		// determine whether the canonical token is bridged or not on this chain
		bridgedAddress, err = p.destERC20Vault.CanonicalToBridged(opts, chainID, addr)
	}

	if eventType == relayer.EventTypeSendERC721 && destChainID.Cmp(chainID) != 0 {
		// determine whether the canonical token is bridged or not on this chain
		bridgedAddress, err = p.destERC721Vault.CanonicalToBridged(opts, chainID, addr)
	}

	if eventType == relayer.EventTypeSendERC1155 && destChainID.Cmp(chainID) != 0 {
		// determine whether the canonical token is bridged or not on this chain
		bridgedAddress, err = p.destERC1155Vault.CanonicalToBridged(opts, chainID, addr)
	}

	if err != nil {
		return false, err
	}

	return bridgedAddress == relayer.ZeroAddress, nil
}

// hardcodeGasLimit determines a viable gas limit when we can get
// unable to estimate gas for contract deployments within the contract code.
// if we get an error or the gas is 0, lets manual set high gas limit and ignore error,
// and try to actually send.
// if contract has not been deployed, we need much higher gas limit, otherwise, we can
// send lower.
func (p *Processor) hardcodeGasLimit(
	ctx context.Context,
	auth *bind.TransactOpts,
	event *bridge.BridgeMessageSent,
	eventType relayer.EventType,
	canonicalToken relayer.CanonicalToken,
) error {
	var bridgedAddress common.Address

	var err error

	switch eventType {
	case relayer.EventTypeSendETH:
		// eth bridges take much less gas, from 250k to 450k.
		auth.GasLimit = 500000
	case relayer.EventTypeSendERC20:
		// determine whether the canonical token is bridged or not on this chain
		bridgedAddress, err = p.destERC20Vault.CanonicalToBridged(
			nil,
			new(big.Int).SetUint64(canonicalToken.ChainID()),
			canonicalToken.Address(),
		)
		if err != nil {
			return errors.Wrap(err, "p.destERC20Vault.CanonicalToBridged")
		}
	case relayer.EventTypeSendERC721:
		// determine whether the canonical token is bridged or not on this chain
		bridgedAddress, err = p.destERC721Vault.CanonicalToBridged(
			nil,
			new(big.Int).SetUint64(canonicalToken.ChainID()),
			canonicalToken.Address(),
		)
		if err != nil {
			return errors.Wrap(err, "p.destERC721Vault.CanonicalToBridged")
		}
	case relayer.EventTypeSendERC1155:
		// determine whether the canonical token is bridged or not on this chain
		bridgedAddress, err = p.destERC1155Vault.CanonicalToBridged(
			nil,
			new(big.Int).SetUint64(canonicalToken.ChainID()),
			canonicalToken.Address(),
		)
		if err != nil {
			return errors.Wrap(err, "p.destERC1155Vault.CanonicalToBridged")
		}
	default:
		return errors.New("unexpected event type")
	}

	if bridgedAddress == relayer.ZeroAddress {
		// needs large gas limit because it has to deploy an ERC20 contract on destination
		// chain. deploying ERC20 can be 2 mil by itself.
		auth.GasLimit = 3000000
	} else {
		// needs larger than ETH gas limit but not as much as deploying ERC20.
		// takes 450-550k gas after signalRoot refactors.
		auth.GasLimit = 600000
	}

	return nil
}

func (p *Processor) setLatestNonce(nonce uint64) {
	p.destNonce = nonce
}

func (p *Processor) saveMessageStatusChangedEvent(
	ctx context.Context,
	receipt *types.Receipt,
	event *bridge.BridgeMessageSent,
) error {
	bridgeAbi, err := abi.JSON(strings.NewReader(bridge.BridgeABI))
	if err != nil {
		return errors.Wrap(err, "abi.JSON")
	}

	m := make(map[string]interface{})

	for _, log := range receipt.Logs {
		topic := log.Topics[0]
		if topic == bridgeAbi.Events["MessageStatusChanged"].ID {
			err = bridgeAbi.UnpackIntoMap(m, "MessageStatusChanged", log.Data)
			if err != nil {
				return errors.Wrap(err, "abi.UnpackIntoInterface")
			}

			break
		}
	}

	if m["status"] != nil {
		// keep same format as other raw events
		data := fmt.Sprintf(`{"Raw":{"transactionHash": "%v"}}`, receipt.TxHash.Hex())

		_, err = p.eventRepo.Save(ctx, relayer.SaveEventOpts{
			Name:         relayer.EventNameMessageStatusChanged,
			Data:         data,
			ChainID:      new(big.Int).SetUint64(event.Message.DestChainId),
			Status:       relayer.EventStatus(m["status"].(uint8)),
			MsgHash:      common.Hash(event.MsgHash).Hex(),
			MessageOwner: event.Message.Owner.Hex(),
			Event:        relayer.EventNameMessageStatusChanged,
		})
		if err != nil {
			return errors.Wrap(err, "svc.eventRepo.Save")
		}
	}

	return nil
}

func (p *Processor) setGasTipOrPrice(ctx context.Context, auth *bind.TransactOpts) error {
	gasTipCap, err := p.destEthClient.SuggestGasTipCap(ctx)
	if err != nil {
		if IsMaxPriorityFeePerGasNotFoundError(err) {
			auth.GasTipCap = FallbackGasTipCap
		} else {
			gasPrice, err := p.destEthClient.SuggestGasPrice(context.Background())
			if err != nil {
				return errors.Wrap(err, "p.destBridge.SuggestGasPrice")
			}
			auth.GasPrice = gasPrice
		}
	}

	auth.GasTipCap = gasTipCap

	return nil
}

func (p *Processor) getCost(ctx context.Context, auth *bind.TransactOpts) (*big.Int, error) {
	if auth.GasTipCap != nil {
		blk, err := p.destEthClient.BlockByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}

		var baseFee *big.Int

		if p.taikoL2 != nil {
			gasUsed := uint32(blk.GasUsed())
			timeSince := uint64(time.Since(time.Unix(int64(blk.Time()), 0)))
			baseFee, err = p.taikoL2.GetBasefee(&bind.CallOpts{Context: ctx}, timeSince, gasUsed)

			if err != nil {
				return nil, errors.Wrap(err, "p.taikoL2.GetBasefee")
			}
		} else {
			cfg := params.NetworkIDToChainConfigOrDefault(p.destChainId)
			baseFee = eip1559.CalcBaseFee(cfg, blk.Header())
		}

		return new(big.Int).Mul(
			new(big.Int).SetUint64(auth.GasLimit),
			new(big.Int).Add(auth.GasTipCap, baseFee)), nil
	} else {
		return new(big.Int).Mul(auth.GasPrice, new(big.Int).SetUint64(auth.GasLimit)), nil
	}
}
