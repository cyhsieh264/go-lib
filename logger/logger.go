package logger

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logger   *logrus.Logger
	logLevel map[string]logrus.Level
}

var (
	once     sync.Once
	instance *Logger
)

func GetLogger() *Logger {
	once.Do(func() {
		logLevel := map[string]logrus.Level{
			"TRACE": logrus.TraceLevel,
			"DEBUG": logrus.DebugLevel,
			"INFO":  logrus.InfoLevel,
			"WARN":  logrus.WarnLevel,
			"ERROR": logrus.ErrorLevel,
			"FATAL": logrus.FatalLevel,
			"PANIC": logrus.PanicLevel,
		}

		instance = &Logger{logrus.New(), logLevel}
	})

	return instance
}

func (logger *Logger) WithReqId(ctx context.Context) *logrus.Entry {
	entry := logger.Logger.WithContext(ctx)
	entry = entry.WithField("req_id", ctx.Value("X-Request-ID"))
	return entry
}
