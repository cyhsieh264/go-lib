package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var (
	once     sync.Once
	instance *Logger
)

// TODO: set log level and implement custom log formatter

func New() *Logger {
	once.Do(func() {
		instance = &Logger{logrus.New()}
	})

	return instance
}
