package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

func SetLevel(l logrus.Level) {
	logger.SetLevel(l)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Context(ctx context.Context) *logrus.Entry {
	return logger.WithContext(ctx)
}
