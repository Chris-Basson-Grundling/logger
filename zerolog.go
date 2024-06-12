package logger

import (
	"context"
	"github.com/rs/zerolog"
	"os"
)

type ZeroLogger struct {
	logger *zerolog.Logger
	ctx    context.Context
}

func NewZeroLogger(loggerType string, ctx context.Context) *ZeroLogger {
	logger := zerolog.New(os.Stdout)

	return &ZeroLogger{logger: &logger, ctx: ctx}
}

func (l *ZeroLogger) Debug(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Debug().Any("args", fields).Msg(msg)
}

func (l *ZeroLogger) Info(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Info().Any("args", fields).Msg(msg)
}

func (l *ZeroLogger) Warn(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Warn().Any("args", fields).Msg(msg)
}

func (l *ZeroLogger) Error(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Error().Any("args", fields).Msg(msg)
}

func (l *ZeroLogger) Fatal(msg string, fields map[string]interface{}) {
	l.addContextCommonFields(fields)
	l.logger.Fatal().Any("args", fields).Msg(msg)
}

func (l *ZeroLogger) addContextCommonFields(fields map[string]interface{}) {
	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}
