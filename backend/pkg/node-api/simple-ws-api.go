package node_api

import (
	"go.uber.org/zap"
	"net/http"
)

type SimpleWSAPI struct {
	Logger *zap.Logger
}

func (n SimpleWSAPI) Fetch(client *http.Client, url string) (*NodeAPIResult, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Sec-WebSocket-Key", "aGVsbG93b3JsZAo=")

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		n.Logger.Debug("node is not alive, got error", zap.String("url", url), zap.Error(err))
		return &NodeAPIResult{Alive: false, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	} else if res.StatusCode != 101 {
		n.Logger.Debug("node is not alive, status > 500", zap.String("url", url), zap.Int("status", res.StatusCode))
		return &NodeAPIResult{Alive: false, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	} else {
		n.Logger.Debug("node is alive", zap.String("url", url), zap.Int("status", res.StatusCode))
		return &NodeAPIResult{Alive: true, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	}
}
