// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis          redis.RedisConf `json:",inherit"`
	DataSource     string          `json:",inherit"`
	SqlLog         bool            `json:",optional,default=false,inherit"`
	DefaultChainId int64           `json:""`
	PriceCronSpec  string          `json:",default=@every 30m"`
	Terms          []Term          `json:""`
	EvmHost        string          `json:""`
	Auth           struct {
		AccessSecret string
		AccessExpire int64
	}
	UniBTC       string        `json:""`
	ExtraRewards []ExtraReward `json:",optional"`
}

type Term struct {
	Symbol  string
	TermMd5 string //terms content hash. md5
}

type ExtraReward struct {
	Symbol         string `json:""`
	Url            string `json:""`
	BaseProfitStep int64  `json:",default=86400"`
}
