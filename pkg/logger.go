package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AppLog interface {
	Info(msg string)
	Debug(msg string)
	Warning(msg string)
	Error(msg interface{})
}

type appLogsZap struct {
	log *zap.Logger
}

func NewAppLogsZap() AppLog {

	var log *zap.Logger

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &appLogsZap{log: log}
}

func (l *appLogsZap) Info(msg string) {

	l.log.Info(msg)
}

func (l *appLogsZap) Debug(msg string) {

	l.log.Debug(msg)
}

func (l *appLogsZap) Warning(msg string) {

	l.log.Warn(msg)
}

func (l *appLogsZap) Error(msg interface{}) {

	switch v := msg.(type) {
	case error:

		l.log.Error(v.Error())
	case string:

		l.log.Error(v)
	}
}
