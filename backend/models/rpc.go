package models

import (
	"time"
)

type Status string

const (
	StatusActive Status = "active"
	StatusDown   Status = "down"
)

type RPC struct {
	Id   string `json:"id" yaml:"id" bson:"_id"`
	HTTP string `json:"http" yaml:"http" bson:"http"`
	WS   string `json:"ws" yaml:"ws" bson:"ws"`

	Status Status `json:"status" yaml:"Status" bson:"status"`

	ChainId  string `json:"chainId" yaml:"chainId" bson:"chainId"`
	Chain    string `json:"chain" yaml:"chain" bson:"chain"`
	Network  string `json:"network" yaml:"network" bson:"network"`
	Provider string `json:"provider" yaml:"provider" bson:"provider"`

	AddedAt   time.Time `json:"addedAt" yaml:"addedAt" bson:"addedAt"`
	CheckedAt time.Time `json:"checkedAt" yaml:"checkedAt" bson:"checkedAt"`
}

func CompareRpcMetadata(rpc1, rpc2 *RPC) bool {
	return rpc1.ChainId == rpc2.ChainId &&
		rpc1.Chain == rpc2.Chain &&
		rpc1.Network == rpc2.Network &&
		rpc1.Provider == rpc2.Provider
}
