// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"cuniBTCReward/api/internal/config"
	unibtcprice "cuniBTCReward/api/internal/crontab"
	"cuniBTCReward/pkg/gormz"
	"cuniBTCReward/service/crons"
	"time"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config           config.Config
	Database         *gorm.DB
	Cron             *crons.ScanCron
	UniBtcPriceCron  *unibtcprice.CoinGeckoUniBTC
	Redis            *redis.Redis
	SignTermsLimiter *limit.PeriodLimit
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: gormz.NewGormLogger(c.SqlLog),
	})
	logx.Must(err)
	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	rds := redis.MustNewRedis(c.Redis)

	uniBtcPriceCron := unibtcprice.NewCoinGeckoUniBTC(c)
	crontab := crons.New()
	_, err = crontab.AddFunc(c.PriceCronSpec, uniBtcPriceCron.CoinGeckoUniBTCCron)
	logx.Must(err)
	uniBtcPriceCron.CoinGeckoUniBTCCron()
	logx.Infof("add cron priceCron scan spec: %v", c.PriceCronSpec)
	crontab.Run()

	return &ServiceContext{
		Config:           c,
		Database:         db,
		Cron:             crontab,
		UniBtcPriceCron:  uniBtcPriceCron,
		Redis:            rds,
		SignTermsLimiter: limit.NewPeriodLimit(60, 3, rds, "cuniBTC:signTerm:rate:"),
	}
}
