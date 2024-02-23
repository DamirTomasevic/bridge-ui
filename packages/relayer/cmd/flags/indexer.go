package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	SrcBridgeAddress = &cli.StringFlag{
		Name:     "srcBridgeAddress",
		Usage:    "Bridge address on the source chain",
		Required: true,
		Category: indexerCategory,
		EnvVars:  []string{"SRC_BRIDGE_ADDRESS"},
	}
	DestBridgeAddress = &cli.StringFlag{
		Name:     "destBridgeAddress",
		Usage:    "Bridge address for the destination chain",
		Required: true,
		Category: commonCategory,
		EnvVars:  []string{"DEST_BRIDGE_ADDRESS"},
	}
)

// optional
var (
	BlockBatchSize = &cli.Uint64Flag{
		Name:     "blockBatchSize",
		Usage:    "Block batch size when iterating through blocks",
		Value:    10,
		Category: indexerCategory,
		EnvVars:  []string{"BLOCK_BATCH_SIZE"},
	}
	MaxNumGoroutines = &cli.Uint64Flag{
		Name:     "maxNumGoroutines",
		Usage:    "Max number of goroutines to spawn simultaneously when indexing",
		Value:    10,
		Category: indexerCategory,
		EnvVars:  []string{"NUM_GOROUTINES"},
	}
	SubscriptionBackoff = &cli.Uint64Flag{
		Name:     "subscriptionBackoff",
		Usage:    "Subscription backoff in seconds",
		Value:    30,
		Category: indexerCategory,
		EnvVars:  []string{"SUBSCRIPTION_BACKOFF_IN_SECONDS"},
	}
	SyncMode = &cli.StringFlag{
		Name:     "syncMode",
		Usage:    "Mode of syncing. Pass in 'sync' to continue, and 'resync' to start from genesis again.",
		Value:    "sync",
		Category: indexerCategory,
		EnvVars:  []string{"SYNC_MODE"},
	}
	WatchMode = &cli.StringFlag{
		Name: "watchMode",
		Usage: `Mode of watching the chain. Options are:
		filter: only filter the chain, when caught up, exit
		subscribe: do not filter the chain, only subscribe to new events
		filter-and-subscribe: the default behavior, filter the chain and subscribe when caught up
		crawl-past-blocks: crawl past blocks
		`,
		Value:    "filter-and-subscribe",
		Category: indexerCategory,
		EnvVars:  []string{"WATCH_MODE"},
	}
	SrcTaikoAddress = &cli.StringFlag{
		Name:     "srcTaikoAddress",
		Usage:    "Taiko address on the source chain, required if L1=>L2, not if L2=>L1",
		Category: indexerCategory,
		EnvVars:  []string{"SRC_TAIKO_ADDRESS"},
	}
	NumLatestBlocksToIgnoreWhenCrawling = &cli.Uint64Flag{
		Name:     "numLatestBlocksToIgnoreWhenCrawling",
		Usage:    "Number of blocks to ingore when crawling chain, should be higher for L2-L1 indexing due to delay",
		Value:    1000,
		Category: indexerCategory,
		EnvVars:  []string{"NUM_LATEST_BLOCKS_TO_IGNORE_WHEN_CRAWLING"},
	}
	TargetBlockNumber = &cli.Uint64Flag{
		Name:     "targetBlockNumber",
		Usage:    "Specify the target block number to process transactions in",
		Required: false,
		Category: indexerCategory,
		EnvVars:  []string{"TARGET_BLOCK_NUMBER"},
	}
)

var IndexerFlags = MergeFlags(CommonFlags, QueueFlags, []cli.Flag{
	SrcBridgeAddress,
	DestBridgeAddress,
	// optional
	SrcTaikoAddress,
	BlockBatchSize,
	MaxNumGoroutines,
	SubscriptionBackoff,
	SyncMode,
	WatchMode,
	NumLatestBlocksToIgnoreWhenCrawling,
	TargetBlockNumber,
})
