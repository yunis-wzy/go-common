package extend_gorm_logger

import (
	"context"
	"go-common/logs"
	"time"

	"gorm.io/gorm/logger"
)

type Logger struct {
	logger.LogLevel
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	cost := float64(time.Since(begin).Nanoseconds()/1e4) / 100.0
	sql, _ := fc()
	logs.CtxInfo(ctx, "GORM LOG SQL:%s Cost:%.2fms", sql, cost)
}

func (l Logger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	logs.CtxInfo(ctx, "GORM LOG "+msg, data...)
}

func (l Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logs.CtxWarn(ctx, "GORM LOG "+msg, data...)
}

func (l Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	logs.CtxError(ctx, "GORM LOG "+msg, data...)
}
