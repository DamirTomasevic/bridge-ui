package encoding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Proof struct {
	AccountProof []byte `abi:"accountProof"`
	StorageProof []byte `abi:"storageProof"`
}

type BlockHeader struct {
	ParentHash       [32]byte       `abi:"parentHash"`
	OmmersHash       [32]byte       `abi:"ommersHash"`
	Beneficiary      common.Address `abi:"beneficiary"`
	StateRoot        [32]byte       `abi:"stateRoot"`
	TransactionsRoot [32]byte       `abi:"transactionsRoot"`
	ReceiptsRoot     [32]byte       `abi:"receiptsRoot"`
	LogsBloom        [8][32]byte    `abi:"logsBloom"`
	Difficulty       *big.Int       `abi:"difficulty"`
	Height           *big.Int       `abi:"height"`
	GasLimit         uint64         `abi:"gasLimit"`
	GasUsed          uint64         `abi:"gasUsed"`
	Timestamp        uint64         `abi:"timestamp"`
	ExtraData        []byte         `abi:"extraData"`
	MixHash          [32]byte       `abi:"mixHash"`
	Nonce            uint64         `abi:"nonce"`
	BaseFeePerGas    *big.Int       `abi:"baseFeePerGas"`
	WithdrawalsRoot  [32]byte       `abi:"withdrawalsRoot"`
}

type SignalProof struct {
	CrossChainSync common.Address `abi:"crossChainSync"`
	Height         uint64         `abi:"height"`
	StorageProof   []byte         `abi:"storageProof"`
	Hops           []Hop          `abi:"hops"`
}

type Hop struct {
	SignalRootRelay common.Address `abi:"signalRootRelay"`
	SignalRoot      [32]byte       `abi:"signalRoot"`
	StorageProof    []byte         `abi:"storageProof"`
}

var hopComponents = []abi.ArgumentMarshaling{
	{
		Name: "signalRootRelay",
		Type: "address",
	},
	{
		Name: "signalRoot",
		Type: "bytes32",
	},
	{
		Name: "storageProof",
		Type: "bytes",
	},
}

var signalProofT, _ = abi.NewType("tuple", "", []abi.ArgumentMarshaling{
	{
		Name: "crossChainSync",
		Type: "address",
	},
	{
		Name: "height",
		Type: "uint64",
	},
	{
		Name: "storageProof",
		Type: "bytes",
	},
	{
		Name:       "hops",
		Type:       "tuple[]",
		Components: hopComponents,
	},
})
