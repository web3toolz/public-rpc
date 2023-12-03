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
	Id       string `json:"id" bson:"_id"`
	HTTP     string `json:"http" bson:"http"`
	WS       string `json:"ws" bson:"ws"`
	Provider string `json:"provider" bson:"provider"`

	Status Status `json:"status" bson:"status"`

	ChainId string `json:"chainId" bson:"chainId"`
	Chain   string `json:"chain" bson:"chain"`
	Network string `json:"network" bson:"network"`

	AddedAt   time.Time `json:"addedAt" bson:"addedAt"`
	CheckedAt time.Time `json:"checkedAt" bson:"checkedAt"`
}
