package logs

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var APIZapLogger *zap.SugaredLogger
var logIDKeyName = "trace_id"
var defaultLogID = "-"

func newZapLogger(output zapcore.WriteSyncer, level int64) *zap.Logger {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // default print time in "2019-07-25T16:49:46.037+0800"
		EncodeDuration: zapcore.NanosDurationEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encCfg)

	return zap.New(zapcore.NewCore(encoder, output, zap.NewAtomicLevelAt(getLogLevel(level))))
}

func Init(level int64) {
	syncer := zap.CombineWriteSyncers(
		os.Stdout,
	)
	APIZapLogger = newZapLogger(syncer, level).Sugar()
}

// 各种基础日志输出
func CtxDebug(ctx context.Context, msg string, fields ...interface{}) {
	APIZapLogger.WithOptions(zap.Fields(extendFiled(ctx)...)).Debugf(msg, fields...)
}

func CtxInfo(ctx context.Context, msg string, fields ...interface{}) {
	APIZapLogger.WithOptions(zap.Fields(extendFiled(ctx)...)).Infof(msg, fields...)
}

func CtxWarn(ctx context.Context, msg string, fields ...interface{}) {
	APIZapLogger.WithOptions(zap.Fields(extendFiled(ctx)...)).Warnf(msg, fields...)
}

func CtxError(ctx context.Context, msg string, fields ...interface{}) {
	APIZapLogger.WithOptions(zap.Fields(extendFiled(ctx)...)).Errorf(msg, fields...)
}

func CtxFatal(ctx context.Context, msg string, fields ...interface{}) {
	APIZapLogger.WithOptions(zap.Fields(extendFiled(ctx)...)).Fatalf(msg, fields...)
}

func extendFiled(ctx context.Context) []zap.Field {
	return []zap.Field{zap.String(logIDKeyName, getLogid(ctx))}
}

func getLogLevel(level int64) zapcore.Level {
	res, ok := levelMap[level]
	if !ok {
		return zap.InfoLevel
	}
	return res
}

func getLogid(ctx context.Context) string {
	if trace.SpanFromContext(ctx).SpanContext().IsValid() {
		return trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	}
	return defaultLogID
}
