package proposer

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/utils"
)

var (
	l1Endpoint      = os.Getenv("L1_NODE_WS_ENDPOINT")
	l2Endpoint      = os.Getenv("L2_EXECUTION_ENGINE_HTTP_ENDPOINT")
	taikoL1         = os.Getenv("TAIKO_L1_ADDRESS")
	taikoL2         = os.Getenv("TAIKO_L2_ADDRESS")
	taikoToken      = os.Getenv("TAIKO_TOKEN_ADDRESS")
	proverEndpoints = "http://localhost:9876,http://localhost:1234"
	tierFee         = 100.0
	proposeInterval = "10s"
	rpcTimeout      = "5s"
)

func (s *ProposerTestSuite) TestNewConfigFromCliContext() {
	goldenTouchAddress, err := s.RPCClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	app := s.SetupApp()

	app.Action = func(cliCtx *cli.Context) error {
		c, err := NewConfigFromCliContext(cliCtx)
		s.Nil(err)
		s.Equal(l1Endpoint, c.L1Endpoint)
		s.Equal(l2Endpoint, c.L2Endpoint)
		s.Equal(taikoL1, c.TaikoL1Address.String())
		s.Equal(taikoL2, c.TaikoL2Address.String())
		s.Equal(taikoToken, c.TaikoTokenAddress.String())
		s.Equal(goldenTouchAddress, crypto.PubkeyToAddress(c.L1ProposerPrivKey.PublicKey))
		s.Equal(goldenTouchAddress, c.L2SuggestedFeeRecipient)
		s.Equal(float64(10), c.ProposeInterval.Seconds())
		s.Equal(1, len(c.LocalAddresses))
		s.Equal(goldenTouchAddress, c.LocalAddresses[0])
		s.Equal(5*time.Second, c.Timeout)
		tierFeeGWei, err := utils.GWeiToWei(tierFee)
		s.Nil(err)
		s.Equal(tierFeeGWei.Uint64(), c.OptimisticTierFee.Uint64())
		s.Equal(tierFeeGWei.Uint64(), c.SgxTierFee.Uint64())
		s.Equal(uint64(15), c.TierFeePriceBump.Uint64())
		s.Equal(uint64(5), c.MaxTierFeePriceBumps)
		s.Equal(true, c.IncludeParentMetaHash)

		for i, e := range strings.Split(proverEndpoints, ",") {
			s.Equal(c.ProverEndpoints[i].String(), e)
		}

		s.Nil(new(Proposer).InitFromCli(context.Background(), cliCtx))
		return nil
	}

	s.Nil(app.Run([]string{
		"TestNewConfigFromCliContext",
		"--" + flags.L1WSEndpoint.Name, l1Endpoint,
		"--" + flags.L2HTTPEndpoint.Name, l2Endpoint,
		"--" + flags.TaikoL1Address.Name, taikoL1,
		"--" + flags.TaikoL2Address.Name, taikoL2,
		"--" + flags.TaikoTokenAddress.Name, taikoToken,
		"--" + flags.L1ProposerPrivKey.Name, encoding.GoldenTouchPrivKey,
		"--" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.TxPoolLocals.Name, goldenTouchAddress.Hex(),
		"--" + flags.RPCTimeout.Name, rpcTimeout,
		"--" + flags.TxGasLimit.Name, "100000",
		"--" + flags.ProverEndpoints.Name, proverEndpoints,
		"--" + flags.OptimisticTierFee.Name, fmt.Sprint(tierFee),
		"--" + flags.SgxTierFee.Name, fmt.Sprint(tierFee),
		"--" + flags.TierFeePriceBump.Name, "15",
		"--" + flags.MaxTierFeePriceBumps.Name, "5",
		"--" + flags.ProposeBlockIncludeParentMetaHash.Name, "true",
	}))
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextPrivKeyErr() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextPrivKeyErr",
		"--" + flags.L1ProposerPrivKey.Name, string(common.FromHex("0x")),
	}), "invalid L1 proposer private key")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextL2RecipErr() {
	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextL2RecipErr",
		"--" + flags.L1ProposerPrivKey.Name, encoding.GoldenTouchPrivKey,
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.MinProposingInternal.Name, proposeInterval,
		"--" + flags.L2SuggestedFeeRecipient.Name, "notAnAddress",
	}), "invalid L2 suggested fee recipient address")
}

func (s *ProposerTestSuite) TestNewConfigFromCliContextTxPoolLocalsErr() {
	goldenTouchAddress, err := s.RPCClient.TaikoL2.GOLDENTOUCHADDRESS(nil)
	s.Nil(err)

	app := s.SetupApp()

	s.ErrorContains(app.Run([]string{
		"TestNewConfigFromCliContextTxPoolLocalsErr",
		"--" + flags.L1ProposerPrivKey.Name, encoding.GoldenTouchPrivKey,
		"--" + flags.ProposeInterval.Name, proposeInterval,
		"--" + flags.MinProposingInternal.Name, proposeInterval,
		"--" + flags.L2SuggestedFeeRecipient.Name, goldenTouchAddress.Hex(),
		"--" + flags.TxPoolLocals.Name, "notAnAddress",
	}), "invalid account in --txpool.locals")
}

func (s *ProposerTestSuite) SetupApp() *cli.App {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: flags.L1WSEndpoint.Name},
		&cli.StringFlag{Name: flags.L2HTTPEndpoint.Name},
		&cli.StringFlag{Name: flags.TaikoL1Address.Name},
		&cli.StringFlag{Name: flags.TaikoL2Address.Name},
		&cli.StringFlag{Name: flags.TaikoTokenAddress.Name},
		&cli.StringFlag{Name: flags.L1ProposerPrivKey.Name},
		&cli.StringFlag{Name: flags.L2SuggestedFeeRecipient.Name},
		&cli.DurationFlag{Name: flags.MinProposingInternal.Name},
		&cli.DurationFlag{Name: flags.ProposeInterval.Name},
		&cli.StringFlag{Name: flags.TxPoolLocals.Name},
		&cli.StringFlag{Name: flags.ProverEndpoints.Name},
		&cli.Uint64Flag{Name: flags.OptimisticTierFee.Name},
		&cli.Uint64Flag{Name: flags.SgxTierFee.Name},
		&cli.DurationFlag{Name: flags.RPCTimeout.Name},
		&cli.Uint64Flag{Name: flags.TierFeePriceBump.Name},
		&cli.Uint64Flag{Name: flags.MaxTierFeePriceBumps.Name},
		&cli.BoolFlag{Name: flags.ProposeBlockIncludeParentMetaHash.Name},
		&cli.StringFlag{Name: flags.AssignmentHookAddress.Name},
	}
	app.Flags = append(app.Flags, flags.TxmgrFlags...)
	app.Action = func(ctx *cli.Context) error {
		_, err := NewConfigFromCliContext(ctx)
		return err
	}
	return app
}
