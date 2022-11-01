package logger

import "go.uber.org/zap"

var (
	Log *zap.Logger
)

func init() {
	var err error
	Log, err = zap.NewProduction()
	if err != nil {
		panic("failed to initialize zap logger..." + err.Error())
	}
}
