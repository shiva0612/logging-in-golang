package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	var err error
	config := zap.NewProductionConfig()

	encoderconfig := zap.NewProductionEncoderConfig()
	encoderconfig.TimeKey = "Timestamp"
	encoderconfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderconfig.StacktraceKey = "" //not printing stacktrace

	config.EncoderConfig = encoderconfig
	log, err = config.Build(zap.AddCallerSkip(1))

	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic("failed to initialize zap logger..." + err.Error())
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}
