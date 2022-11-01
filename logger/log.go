package logger

import "go.uber.org/zap"

var (
	log *zap.Logger
)

func init() {
	var err error
	log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic("failed to initialize zap logger..." + err.Error())
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}
