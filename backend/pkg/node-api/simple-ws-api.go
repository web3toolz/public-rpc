package node_api

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	urllib "net/url"
)

type SimpleWSAPI struct {
	Logger *zap.Logger
}

func (n SimpleWSAPI) Fetch(client *http.Client, url string) (*NodeAPIResult, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	u, err := urllib.Parse(url)

	if err != nil {
		return nil, err
	}

	httpUrl := fmt.Sprintf("https://%s%s?%s", u.Host, u.Path, u.RawQuery)
	req, err := http.NewRequest("GET", httpUrl, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Sec-WebSocket-Key", "N509hQQwSFGfanhbxP3F6g==")
	req.Header.Set("Host", u.Host)
	req.Header.Set("Origin", fmt.Sprintf("https://%s", u.Host))

	req = req.WithContext(ctx)

	res, err := client.Do(req)

	if err != nil {
		n.Logger.Debug("node is not alive, got error", zap.String("url", url), zap.Error(err))
		return &NodeAPIResult{Alive: false, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	} else if res.StatusCode >= 400 {
		n.Logger.Debug("node is not alive, status > 400", zap.String("url", url), zap.Int("status", res.StatusCode))
		return &NodeAPIResult{Alive: false, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	} else {
		n.Logger.Debug("node is alive", zap.String("url", url), zap.Int("status", res.StatusCode))
		return &NodeAPIResult{Alive: true, LatestBlock: 0, LatestBlockTimestamp: 0}, nil
	}
}
