package logger

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logger *logrus.Logger
}

var (
	once     sync.Once
	instance *Logger
)

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{logrus.New()}
	})

	return instance
}

func (logger *Logger) WithReqId(ctx context.Context) *logrus.Entry {
	entry := logger.Logger.WithContext(ctx)
	entry = entry.WithField("req_id", ctx.Value("X-Request-ID"))
	return entry
}
