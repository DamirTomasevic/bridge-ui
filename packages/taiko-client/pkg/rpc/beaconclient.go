package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v4/api/client"
	"github.com/prysmaticlabs/prysm/v4/api/client/beacon"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/blob"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/config"
)

var (
	// Request urls.
	sidecarsRequestURL = "/eth/v1/beacon/blob_sidecars/%d"
	genesisRequestURL  = "/eth/v1/beacon/genesis"
	getConfigSpecPath  = "/eth/v1/config/spec"
)

type ConfigSpec struct {
	SecondsPerSlot string `json:"SECONDS_PER_SLOT"`
}

type GenesisResponse struct {
	Data struct {
		GenesisTime string `json:"genesis_time"`
	} `json:"data"`
}

type BeaconClient struct {
	*beacon.Client

	timeout        time.Duration
	genesisTime    uint64
	secondsPerSlot uint64
}

// NewBeaconClient returns a new beacon client.
func NewBeaconClient(endpoint string, timeout time.Duration) (*BeaconClient, error) {
	cli, err := beacon.NewClient(strings.TrimSuffix(endpoint, "/"), client.WithTimeout(timeout))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Get the genesis time.
	var genesisDetail *GenesisResponse
	resBytes, err := cli.Get(ctx, cli.BaseURL().Path+genesisRequestURL)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resBytes, &genesisDetail); err != nil {
		return nil, err
	}

	genesisTime, err := strconv.Atoi(genesisDetail.Data.GenesisTime)
	if err != nil {
		return nil, err
	}

	log.Info("L1 genesis time", "time", genesisTime)

	// Get the seconds per slot.
	spec, err := getConfigSpec(ctx, cli)
	if err != nil {
		return nil, err
	}

	secondsPerSlot, err := strconv.Atoi(spec.Data.(map[string]interface{})["SECONDS_PER_SLOT"].(string))
	if err != nil {
		return nil, err
	}

	log.Info("L1 seconds per slot", "seconds", secondsPerSlot)

	return &BeaconClient{cli, timeout, uint64(genesisTime), uint64(secondsPerSlot)}, nil
}

// GetBlobs returns the sidecars for a given slot.
func (c *BeaconClient) GetBlobs(ctx context.Context, time uint64) ([]*blob.Sidecar, error) {
	ctxWithTimeout, cancel := ctxWithTimeoutOrDefault(ctx, c.timeout)
	defer cancel()

	slot, err := c.timeToSlot(time)
	if err != nil {
		return nil, err
	}

	var sidecars *blob.SidecarsResponse
	resBytes, err := c.Get(ctxWithTimeout, c.BaseURL().Path+fmt.Sprintf(sidecarsRequestURL, slot))
	if err != nil {
		return nil, err
	}

	return sidecars.Data, json.Unmarshal(resBytes, &sidecars)
}

// timeToSlot returns the slots of the given timestamp.
func (c *BeaconClient) timeToSlot(timestamp uint64) (uint64, error) {
	if timestamp < c.genesisTime {
		return 0, fmt.Errorf("provided timestamp (%v) precedes genesis time (%v)", timestamp, c.genesisTime)
	}
	return (timestamp - c.genesisTime) / c.secondsPerSlot, nil
}

// getConfigSpec retrieve the current configs of the network used by the beacon node.
func getConfigSpec(ctx context.Context, c *beacon.Client) (*config.GetSpecResponse, error) {
	body, err := c.Get(ctx, c.BaseURL().Path+getConfigSpecPath)
	if err != nil {
		return nil, errors.Wrap(err, "error requesting configSpecPath")
	}
	fsr := &config.GetSpecResponse{}
	err = json.Unmarshal(body, fsr)
	if err != nil {
		return nil, err
	}
	return fsr, nil
}
