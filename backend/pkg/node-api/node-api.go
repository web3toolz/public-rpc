package node_api

import (
	"fmt"
)

type NodeAPIResult struct {
	LatestBlock          int64
	LatestBlockTimestamp int64
}

type INodeAPI interface {
	Fetch() (*NodeAPIResult, error)
}

func NewNodeApiFromChain(chain string) (*INodeAPI, error) {
	var nodeApi INodeAPI

	if ChainIsEVM(chain) {
		nodeApi = EVMNodeAPI{}
	} else if ChainIsStarknet(chain) {
		nodeApi = StarknetNodeAPI{}
	} else if ChainIsSolana(chain) {
		nodeApi = SolanaNodeAPI{}
	}

	if nodeApi != nil {
		return &nodeApi, nil
	} else {
		return nil, fmt.Errorf("chain %s is not supported", chain)
	}
}
