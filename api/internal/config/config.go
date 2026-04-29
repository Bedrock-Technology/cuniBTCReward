// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DataSource     string `json:",inherit"`
	SqlLog         bool   `json:",optional,default=false,inherit"`
	DefaultChainId int64  `json:""`
	PriceCronSpec  string `json:",default=@every 30m"`
	CoinGecoKey    string `json:""`
}
