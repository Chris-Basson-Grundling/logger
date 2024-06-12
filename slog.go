package logger

import (
	"context"
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
	ctx    context.Context
}

func NewSlogLogger(loggerType string, ctx context.Context) *SlogLogger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &SlogLogger{logger: logger, ctx: ctx}
}

func (l *SlogLogger) Debug(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Debug(msg, "args", fields)
}

func (l *SlogLogger) Info(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Info(msg, "args", fields)
}

func (l *SlogLogger) Warn(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Warn(msg, "args", fields)
}

func (l *SlogLogger) Error(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Error(msg, "args", fields)
}

func (l *SlogLogger) Fatal(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Error(msg, "args", fields)
}

func (l *SlogLogger) addContextCommonFields(fields map[string]interface{}) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}
