package worker

import (
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"public-rpc/internal/adapters/storage"
	"public-rpc/internal/config"
	"public-rpc/models"
	nodeapi "public-rpc/pkg/node-api"
	"public-rpc/pkg/semaphore"
	"time"
)

type WorkerComponent struct {
	Cfg       config.WorkerConfig
	Logger    *zap.Logger
	Storage   *storage.Storage
	semaphore *semaphore.Semaphore
}

func NewWorkerComponent(cfg config.WorkerConfig, logger *zap.Logger, storage *storage.Storage) (*WorkerComponent, error) {
	return &WorkerComponent{
		Cfg:       cfg,
		Logger:    logger,
		Storage:   storage,
		semaphore: semaphore.NewSemaphore(cfg.Concurrency),
	}, nil
}

func (c *WorkerComponent) RunJob(client *http.Client, rpc models.RPC, waitTime time.Duration) {
	c.Logger.Debug("job started", zap.String("rpc", rpc.Id))
	c.Logger.Debug("sleeping", zap.String("rpc", rpc.Id), zap.Duration("duration", waitTime))
	time.Sleep(waitTime)

	var url string
	var rpcStatus models.Status

	if rpc.HTTP != "" {
		url = rpc.HTTP
	} else if rpc.WS != "" {
		url = rpc.WS
	} else {
		c.Logger.Error("rpc has no http or ws", zap.String("rpc", rpc.Id))
		return
	}

	nodeApi, err := nodeapi.NewNodeApiFromUrl(c.Logger, rpc.HTTP, rpc.WS)

	if err != nil {
		c.Logger.Error("failed to create node api", zap.Error(err))
		return
	}

	result, err := nodeApi.Fetch(client, url)

	if err != nil {
		c.Logger.Error("failed to fetch node", zap.Error(err), zap.Any("rpc", rpc))
		return
	}

	if result.Alive {
		rpcStatus = models.StatusActive
	} else {
		rpcStatus = models.StatusDown
	}

	rpc.Status = rpcStatus
	rpc.CheckedAt = time.Now()

	storage_ := *c.Storage

	err = storage_.UpdateRPC(rpc)

	if err != nil {
		c.Logger.Error("failed to update rpc", zap.Error(err))
		return
	}

	c.Logger.Debug("job finished", zap.Any("rpc", rpc))
}

func (c *WorkerComponent) ScheduleJobs() {
	c.Logger.Info("starting new iterator of worker")
	storage_ := *c.Storage

	rpcList, err := storage_.ListRPC()

	if err != nil {
		c.Logger.Error("failed to fetch rpc", zap.Error(err))
		return
	}

	for _, rpc := range rpcList {
		if rpc.HTTP != "" {
			continue
		}
		c.semaphore.Acquire()
		rpc := rpc
		go func() {
			// get random sleep time from 0 to 10 secs
			waitTime := rand.Intn(10000)
			client := http.Client{Timeout: time.Second * 5}
			c.RunJob(&client, rpc, time.Millisecond*time.Duration(waitTime))
			c.semaphore.Release()
		}()
	}

	c.Logger.Debug("jobs scheduled", zap.Int("count", len(rpcList)))
}

func (c *WorkerComponent) Run() error {
	ticker := time.NewTicker(c.Cfg.Interval)

	for {
		select {
		case <-ticker.C:
			go c.ScheduleJobs()
		}
	}
}
