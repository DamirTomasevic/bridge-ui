package processor

import (
	"context"

	"log/slog"

	"github.com/pkg/errors"
	"github.com/taikoxyz/taiko-mono/packages/relayer"
)

var (
	errImpossible = errors.New("impossible to process")
)

// isProfitable determines whether a message is profitable or not. It should
// check the processing fee, if one does not exist at all, it is definitely not
// profitable. Otherwise, we compare it to the estimated cost.
func (p *Processor) isProfitable(
	ctx context.Context,
	fee uint64,
	gasLimit uint64,
	destChainBaseFee uint64,
	gasTipCap uint64,
) (bool, error) {
	var shouldProcess bool = false

	if fee == 0 || gasLimit == 0 {
		slog.Info("unprofitable: no gasLimit or processingFee",
			"processingFee", fee,
			"gasLimit", gasLimit,
		)

		return shouldProcess, errImpossible
	}

	// if processing fee is higher than baseFee * gasLimit,
	// we should process.
	estimatedOnchainFee := (destChainBaseFee + gasTipCap) * uint64(gasLimit)
	if fee > estimatedOnchainFee {
		shouldProcess = true
	}

	slog.Info("isProfitable",
		"processingFee", fee,
		"destChainBaseFee", destChainBaseFee,
		"gasLimit", gasLimit,
		"shouldProcess", shouldProcess,
		"estimatedOnchainFee", estimatedOnchainFee,
	)

	if !shouldProcess {
		relayer.UnprofitableMessagesDetected.Inc()

		return false, nil
	}

	return true, nil
}
