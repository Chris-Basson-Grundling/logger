package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

func NewLogrusLogger(loggerType string, ctx context.Context) *LogrusLogger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Level = logrus.DebugLevel

	return &LogrusLogger{logger: logger, ctx: ctx}
}

func (l *LogrusLogger) Debug(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.WithFields(fields).Debug(msg)
}

func (l *LogrusLogger) Info(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.WithFields(fields).Info(msg)
}

func (l *LogrusLogger) Warn(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.WithFields(fields).Warn(msg)
}

func (l *LogrusLogger) Error(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.WithFields(fields).Error(msg)
}

func (l *LogrusLogger) Fatal(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.WithFields(fields).Fatal(msg)
}

func (l *LogrusLogger) addContextCommonFields(fields map[string]interface{}) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}
