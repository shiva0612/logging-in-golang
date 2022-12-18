package main

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.Encoding = "console"
	zlog, _ := config.Build(zap.AddStacktrace(zap.ErrorLevel)) //prints stacktrace only for error level and above
	log = zlog.Sugar()
}

func main() {

	if err := print_log_at_source(); err != nil {
		return
	}
}

func print_log_at_source() error {

	log.Infoln("in print_log_at_source")
	err := call1()
	if err != nil {
		return err
	}
	return nil
}

func call1() error {

	log.Infoln("in call1")
	err := call2()
	if err != nil {
		return err
	}
	return nil
}

func call2() error {
	log.Infof("in call2")

	log.Errorf("im error in call2")
	return errors.New("call2 error")
}
