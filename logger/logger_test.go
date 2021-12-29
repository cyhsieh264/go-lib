package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	log := New()
	log.Info("test log info")
	log.Warn("test log warn")
	log.Error("test log error")
	log.Errorf("%s", "log format string")
	log.WithField("k", "v").Error("log with field")
}
