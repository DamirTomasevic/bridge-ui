package main

import (
	"flag"

	"github.com/taikoxyz/taiko-mono/packages/eventindexer"
	"github.com/taikoxyz/taiko-mono/packages/eventindexer/cli"
)

func main() {
	modePtr := flag.String("mode", string(eventindexer.SyncMode), `mode to run in. 
	options:
	  sync: continue syncing from previous block
	  resync: restart syncing from block 0
	  fromBlock: restart syncing from specified block number
	`)

	watchModePtr := flag.String("watch-mode", string(eventindexer.FilterAndSubscribeWatchMode), `watch mode to run in. 
	options:
	  filter: only filter previous messages
	  subscribe: only subscribe to new messages
	  filter-and-subscribe: catch up on all previous messages, then subscribe to new messages
	`)

	flag.Parse()

	cli.Run(
		eventindexer.Mode(*modePtr),
		eventindexer.WatchMode(*watchModePtr),
	)
}
