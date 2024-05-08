package txlistdecoder

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg"
)

// CalldataFetcher is responsible for fetching the txList bytes from the transaction's calldata.
type CalldataFetcher struct{}

// NewCalldataTxListFetcher creates a new CalldataFetcher instance.
func (d *CalldataFetcher) Fetch(
	_ context.Context,
	tx *types.Transaction,
	meta *bindings.TaikoDataBlockMetadata,
) ([]byte, error) {
	if meta.BlobUsed {
		return nil, pkg.ErrBlobUsed
	}

	return encoding.UnpackTxListBytes(tx.Data())
}
