package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var levelMap = map[int64]zapcore.Level{
	LevelDebug: zap.DebugLevel,
	LevelInfo:  zap.InfoLevel,
	LevelWarn:  zap.WarnLevel,
	LevelError: zap.ErrorLevel,
	LevelFatal: zap.FatalLevel,
}
