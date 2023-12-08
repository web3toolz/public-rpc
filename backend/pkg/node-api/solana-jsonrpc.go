package node_api

import "net/http"

type SolanaNodeAPI struct {
}

func (n SolanaNodeAPI) Fetch(client *http.Client, url string) (*NodeAPIResult, error) {
	return &NodeAPIResult{
		LatestBlock:          0,
		LatestBlockTimestamp: 0,
	}, nil
}

func ChainIsSolana(chain string) bool {
	return chain == "solana"
}
