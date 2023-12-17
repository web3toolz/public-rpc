package node_api

import (
	"go.uber.org/zap"
	"net/http"
)

type SimpleHTTPAPI struct {
	Logger *zap.Logger
}

func (s SimpleHTTPAPI) Fetch(client *http.Client, url string) (*NodeAPIResult, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		s.Logger.Debug("node is not alive, got error", zap.String("url", url), zap.Error(err))
		return &NodeAPIResult{Alive: false, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	} else if res.StatusCode > 400 {
		s.Logger.Debug("node is not alive, status > 400", zap.String("url", url), zap.Int("status", res.StatusCode))
		return &NodeAPIResult{Alive: false, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	} else {
		s.Logger.Debug("node is alive", zap.String("url", url), zap.Int("status", res.StatusCode))
		return &NodeAPIResult{Alive: true, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	}
}
