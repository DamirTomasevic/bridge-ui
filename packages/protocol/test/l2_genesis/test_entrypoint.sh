#!/bin/bash

set -eou pipefail

DIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)

if ! command -v docker &> /dev/null 2>&1; then
    echo "ERROR: `docker` command not found"
    exit 1
fi

if ! docker info > /dev/null 2>&1; then
    echo "ERROR: docker daemon isn't running"
    exit 1
fi

GENESIS_JSON=$(cd "$(dirname "$DIR/../../..")"; pwd)/deployments/l2_genesis.json
TESTNET_CONFIG=$DIR/testnet/docker-compose.yml

touch $GENESIS_JSON

echo '
{
  "config": {
    "chainId": 1337,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "muirGlacierBlock": 0,
    "berlinBlock": 0,
    "londonBlock": 0,
    "arrowGlacierBlock": 0,
    "clique": {
      "period": 0,
      "epoch": 30000
    }
  },
  "gasLimit": "10000000",
  "difficulty": "1",
  "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000df08f82de32b8d460adbe8d72043e3a7e25a3b390000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "alloc":
' > $GENESIS_JSON

echo "Starting generate_l2_genesis tests..."

# compile the contracts to get latest bytecode
yarn clean && yarn compile

# run the task
yarn run generate:genesis $DIR/test_config.json

# generate complete genesis json
cat $DIR/../../deployments/l2_genesis_alloc.json >> $GENESIS_JSON

echo '}' >> $GENESIS_JSON

# start a geth instance and init with the output genesis json
echo ""
echo "Start docker compose network..."

docker compose -f $TESTNET_CONFIG down &> /dev/null
docker compose -f $TESTNET_CONFIG up -d

echo ""
echo "Start testing..."

TEST_L2_GENESIS=true npx hardhat test --grep "Generate L2 Genesis"

docker compose -f $TESTNET_CONFIG down
