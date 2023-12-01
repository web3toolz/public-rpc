package node_api

type SolanaNodeAPI struct {
}

func (n SolanaNodeAPI) Fetch() (*NodeAPIResult, error) {
	return &NodeAPIResult{
		LatestBlock:          0,
		LatestBlockTimestamp: 0,
	}, nil
}

func ChainIsSolana(chain string) bool {
	return chain == "solana"
}
