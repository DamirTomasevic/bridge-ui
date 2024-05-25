package prover

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/utils"
)

var (
	l1WsEndpoint   = os.Getenv("L1_NODE_WS_ENDPOINT")
	l1HttpEndpoint = os.Getenv("L1_NODE_HTTP_ENDPOINT")
	l2WsEndpoint   = os.Getenv("L2_EXECUTION_ENGINE_WS_ENDPOINT")
	l2HttpEndpoint = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	l1NodeVersion  = "1.0.0"
	l2NodeVersion  = "0.1.0"
	taikoL1        = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2        = os.Getenv("TAIKO_L2_ADDRESS")
	allowance      = 10.0
	rpcTimeout     = 5 * time.Second
	minTierFee     = 1024.0
)

func (s *ProverTestSuite) TestNewConfigFromCliContextGuardianProver() {
	app := s.SetupApp()
	app.Action = func(ctx *cli.Context) error {
		c, err := NewConfigFromCliContext(ctx)
		s.Nil(err)
		s.Equal(l1WsEndpoint, c.L1WsEndpoint)
		s.Equal(l1HttpEndpoint, c.L1HttpEndpoint)
		s.Equal(l2WsEndpoint, c.L2WsEndpoint)
		s.Equal(l2HttpEndpoint, c.L2HttpEndpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(
			crypto.PubkeyToAddress(s.p.cfg.L1ProverPrivKey.PublicKey),
			crypto.PubkeyToAddress(c.L1ProverPrivKey.PublicKey),
		)
		s.True(c.Dummy)
		s.Equal("", c.Graffiti)
		s.True(c.ProveUnassignedBlocks)
		s.True(c.ContesterMode)
		s.Equal(rpcTimeout, c.RPCTimeout)
		s.Equal(uint64(8), c.Capacity)
		tierFeeGWei, err := utils.GWeiToWei(minTierFee)
		s.Nil(err)
		s.Equal(tierFeeGWei.Uint64(), c.MinOptimisticTierFee.Uint64())
		s.Equal(tierFeeGWei.Uint64(), c.MinSgxTierFee.Uint64())
		s.Equal(c.L1NodeVersion, l1NodeVersion)
		s.Equal(c.L2NodeVersion, l2NodeVersion)
		s.Nil(new(Prover).InitFromCli(context.Background(), ctx))
		s.True(c.ProveUnassignedBlocks)
		s.Equal(uint64(100), c.MaxProposedIn)
		s.Equal(os.Getenv("ASSIGNMENT_HOOK_ADDRESS"), c.AssignmentHookAddress.String())
		allowanceWithDecimal, err := utils.EtherToWei(allowance)
		s.Nil(err)
		s.Equal(allowanceWithDecimal.Uint64(), c.Allowance.Uint64())

		return err
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContextGuardianProver",
		"--" + flags.L1WSEndpoint.Name, l1WsEndpoint,
		"--" + flags.L1HTTPEndpoint.Name, l1HttpEndpoint,
		"--" + flags.L2WSEndpoint.Name, l2WsEndpoint,
		"--" + flags.L2HTTPEndpoint.Name, l2HttpEndpoint,
		"--" + flags.TaikoL1Address.Name, taikoL1,
		"--" + flags.TaikoL2Address.Name, taikoL2,
		"--" + flags.L1ProverPrivKey.Name, os.Getenv("L1_PROVER_PRIVATE_KEY"),
		"--" + flags.StartingBlockID.Name, "0",
		"--" + flags.RPCTimeout.Name, "5s",
		"--" + flags.TxGasLimit.Name, "100000",
		"--" + flags.Dummy.Name,
		"--" + flags.MinOptimisticTierFee.Name, fmt.Sprint(minTierFee),
		"--" + flags.MinSgxTierFee.Name, fmt.Sprint(minTierFee),
		"--" + flags.ProverCapacity.Name, "8",
		"--" + flags.GuardianProverMajority.Name, os.Getenv("GUARDIAN_PROVER_CONTRACT_ADDRESS"),
		"--" + flags.GuardianProverMinority.Name, os.Getenv("GUARDIAN_PROVER_MINORITY_ADDRESS"),
		"--" + flags.AssignmentHookAddress.Name, os.Getenv("ASSIGNMENT_HOOK_ADDRESS"),
		"--" + flags.Graffiti.Name, "",
		"--" + flags.ProveUnassignedBlocks.Name,
		"--" + flags.MaxProposedIn.Name, "100",
		"--" + flags.Allowance.Name, fmt.Sprint(allowance),
		"--" + flags.L1NodeVersion.Name, l1NodeVersion,
		"--" + flags.L2NodeVersion.Name, l2NodeVersion,
		"--" + flags.RaikoHostEndpoint.Name, "https://dummy.raiko.xyz",
	}))
}

func (s *ProverTestSuite) TestNewConfigFromCliContextProverKeyError() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1ProverPrivKey.Name, "0x",
	}), "invalid L1 prover private key")
}

func (s *ProverTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L1HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.L2WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.L1ProverPrivKey.Name},
		&cli.Uint64Flag{Name: flags.StartingBlockID.Name},
		&cli.BoolFlag{Name: flags.Dummy.Name},
		&cli.StringFlag{Name: flags.GuardianProverMajority.Name},
		&cli.StringFlag{Name: flags.GuardianProverMinority.Name},
		&cli.StringFlag{Name: flags.Graffiti.Name},
		&cli.BoolFlag{Name: flags.ProveUnassignedBlocks.Name},
		&cli.DurationFlag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.ProverCapacity.Name},
		&cli.Uint64Flag{Name: flags.MinOptimisticTierFee.Name},
		&cli.Uint64Flag{Name: flags.MinSgxTierFee.Name},
		&cli.Uint64Flag{Name: flags.MaxProposedIn.Name},
		&cli.StringFlag{Name: flags.AssignmentHookAddress.Name},
		&cli.StringFlag{Name: flags.Allowance.Name},
		&cli.StringFlag{Name: flags.ContesterMode.Name},
		&cli.StringFlag{Name: flags.L1NodeVersion.Name},
		&cli.StringFlag{Name: flags.L2NodeVersion.Name},
		&cli.StringFlag{Name: flags.RaikoHostEndpoint.Name},
	}
	app.Flags = append(app.Flags, flags.TxmgrFlags...)
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		s.NotNil(err)
		return err
	}
	return app
}
