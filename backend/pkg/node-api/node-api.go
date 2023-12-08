package node_api

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type NodeAPIResult struct {
	Alive                bool
	LatestBlock          int64
	LatestBlockTimestamp int64
}

type INodeAPI interface {
	Fetch(client *http.Client, url string) (*NodeAPIResult, error)
}

func NewNodeApiFromUrl(logger *zap.Logger, http string, ws string) (INodeAPI, error) {
	var nodeApi INodeAPI

	if http != "" {
		nodeApi = SimpleHTTPAPI{Logger: logger}
	} else if ws != "" {
		nodeApi = SimpleWSAPI{Logger: logger}
	}

	if nodeApi != nil {
		return nodeApi, nil
	} else {
		return nil, fmt.Errorf("no node api found")
	}
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
