package node_api

import "net/http"

type StarknetNodeAPI struct {
}

func (n StarknetNodeAPI) Fetch(client *http.Client, url string) (*NodeAPIResult, error) {
	return &NodeAPIResult{
		LatestBlock:          0,
		LatestBlockTimestamp: 0,
	}, nil
}

func ChainIsStarknet(chain string) bool {
	return chain == "starknet"
}
