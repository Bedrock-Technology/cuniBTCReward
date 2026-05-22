// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"cuniBTCReward/api/internal/config"
	unibtcprice "cuniBTCReward/api/internal/crontab"
	"cuniBTCReward/pkg/gormz"
	"cuniBTCReward/service/crons"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	Database        *gorm.DB
	Cron            *crons.ScanCron
	UniBtcPriceCron *unibtcprice.CoinGeckoUniBTC
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: gormz.NewGormLogger(c.SqlLog),
	})
	logx.Must(err)

	uniBtcPriceCron := unibtcprice.NewCoinGeckoUniBTC(c)
	crontab := crons.New()
	_, err = crontab.AddFunc(c.PriceCronSpec, uniBtcPriceCron.CoinGeckoUniBTCCron)
	logx.Must(err)
	uniBtcPriceCron.CoinGeckoUniBTCCron()
	logx.Infof("add cron priceCron scan spec: %v", c.PriceCronSpec)
	crontab.Run()

	return &ServiceContext{
		Config:          c,
		Database:        db,
		Cron:            crontab,
		UniBtcPriceCron: uniBtcPriceCron,
	}
}
