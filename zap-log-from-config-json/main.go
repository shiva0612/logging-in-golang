package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func main() {

	level := flag.String("level", "info", "options availabel = info,debug,warn,error,fatal,panic")
	flag.Parse()

	var err error
	logger, err = init_logger("logconfig.json", "app.log", *level)
	if err != nil {
		panic("could not configure zap logger: " + err.Error())
	}

	// logger.Info("zap log")
	slog := logger.Sugar()
	slog.Infof("sugar log")
	slog.Warnf("sugar warning")
	slog.Panicf("sugra panic")
}

func init_logger(configName string, logfileName string, level string) (*zap.Logger, error) {

	b, _ := ioutil.ReadFile(configName)

	zap_config := zap.NewProductionConfig()
	json.Unmarshal(b, &zap_config)

	zap_config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	zap_level, _ := zapcore.ParseLevel(level)
	zap_config.Level = zap.NewAtomicLevelAt(zap_level)

	return zap_config.Build(zap.AddStacktrace(zap.ErrorLevel))
}
