package node_api

type EVMNodeAPI struct {
}

func (n EVMNodeAPI) Fetch() (*NodeAPIResult, error) {
	return &NodeAPIResult{
		LatestBlock:          0,
		LatestBlockTimestamp: 0,
	}, nil
}

func ChainIsEVM(chain string) bool {
	evmChains := []string{
		"ethereum",
		"ethereumclassic",
		"xdai",
		"poa",
		"binance",
		"bsc",
		"bnb",
		"optimism",
		"arbitrum",
		"polygon",
		"avalanche",
		"celo",
		"fuse",
		"heco",
		"okex",
		"thundercore",
		"tomochain",
	}

	for _, evmChain := range evmChains {
		if evmChain == chain {
			return true
		}
	}

	return false
}
