package gormz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	slowThreshold time.Duration
	log           logx.Logger
}

func NewGormLogger() *GormLogger {
	return &GormLogger{
		slowThreshold: 200 * time.Millisecond,
		log:           logx.WithCallerSkip(3),
	}
}

var _ logger.Interface = (*GormLogger)(nil)

func (g *GormLogger) LogMode(lev logger.LogLevel) logger.Interface {
	logx.SetLevel(uint32(lev))
	return g
}

func (g *GormLogger) Info(ctx context.Context, msg string, data ...any) {
	g.log.WithContext(ctx).Infof(msg, data)
}

func (g *GormLogger) Warn(ctx context.Context, msg string, data ...any) {
	g.log.WithContext(ctx).Errorf(msg, data)
}

func (g *GormLogger) Error(ctx context.Context, msg string, data ...any) {
	g.log.WithContext(ctx).Errorf(msg, data)
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logFields := []logx.LogField{
		logx.Field("sql", sql),
		logx.Field("elapsed", fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)),
		logx.Field("rows", rows),
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			g.log.WithContext(ctx).Infof("%v", logFields)
		} else {
			logFields = append(logFields, logx.Field("catch error", err))
			g.log.WithContext(ctx).Infof("%v", logFields)
		}
	}
	if g.slowThreshold != 0 && elapsed > g.slowThreshold {
		g.log.WithContext(ctx).Slowf("%v", logFields)
	}
	g.log.WithContext(ctx).Infof("%v", logFields)
}
