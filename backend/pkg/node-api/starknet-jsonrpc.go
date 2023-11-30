package node_api

type StarknetNodeAPI struct {
}

func (n StarknetNodeAPI) Fetch() (*NodeAPIResult, error) {
	return &NodeAPIResult{
		LatestBlock:          0,
		LatestBlockTimestamp: 0,
	}, nil
}

func ChainIsStarknet(chain string) bool {
	return chain == "starknet"
}
