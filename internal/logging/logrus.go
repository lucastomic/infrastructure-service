package logging

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{logrus.New()}
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}
