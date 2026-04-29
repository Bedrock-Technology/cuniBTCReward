package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	EvmScanConf
}

type EvmScanConf struct {
	service.ServiceConf
	DataSource   string `json:",inherit"`
	SqlLog       bool   `json:",optional,default=false,inherit"`
	ChainInfo    []ChainInfo
	LogsScanSpec string
	NotifySlack  string `json:",optional,inherit"`
	LogSlack     string `json:",optional,inherit"`
}

type EvmClient struct {
	ChainId   uint
	ChainName string
	Host      string
	Request   int
	PeriodSec int
}

type ChainInfo struct {
	Client  EvmClient
	Factory string
}
