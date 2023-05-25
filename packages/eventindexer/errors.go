package eventindexer

import "github.com/cyberhorsey/errors"

var (
	ErrNoEthClient       = errors.Validation.NewWithKeyAndDetail("ERR_NO_ETH_CLIENT", "EthClient is required")
	ErrNoEventRepository = errors.Validation.NewWithKeyAndDetail("ERR_NO_EVENT_REPOSITORY", "EventRepository is required")
	ErrNoStatRepository  = errors.Validation.NewWithKeyAndDetail("ERR_NO_STAT_REPOSITORY", "StatRepository is required")
	ErrNoBlockRepository = errors.Validation.NewWithKeyAndDetail(
		"ERR_NO_BLOCK_REPOSITORY",
		"BlockRepository is required",
	)
	ErrNoCORSOrigins = errors.Validation.NewWithKeyAndDetail("ERR_NO_CORS_ORIGINS", "CORS Origins are required")
	ErrNoRPCClient   = errors.Validation.NewWithKeyAndDetail("ERR_NO_RPC_CLIENT", "RPCClient is required")
	ErrInvalidMode   = errors.Validation.NewWithKeyAndDetail("ERR_INVALID_MODE", "Mode not supported")
)
