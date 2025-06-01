package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
    Logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
    // Configuración personalizada
    config := zap.NewDevelopmentConfig() // Cambia a un formato más legible
    config.EncoderConfig.TimeKey = "time" // Cambia "ts" a "time"
    config.EncoderConfig.CallerKey = "" // Cambia "caller" a "source"
    config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05") // Formato sin milisegundos

    zapLogger, _ := config.Build()
    return &ZapLogger{Logger: zapLogger.Sugar()}
}

func (l *ZapLogger) Info(args ...any) {
    l.Logger.Info(args...)
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...any) {
    l.Logger.Infow(msg, keysAndValues...)}

func (l *ZapLogger) Error(args ...any) {
    l.Logger.Error(args...)
}

func (l *ZapLogger) Debug(args ...any) {
    l.Logger.Debug(args...)
}

func (l *ZapLogger) Warn(args ...any) {
    l.Logger.Warn(args...)
}

func (l *ZapLogger) Errorw(msg string, keysAndValues ...any) {
    l.Logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Warnf(msg string, keysAndValues ...any) {
    l.Logger.Warnw(msg, keysAndValues...)
}
