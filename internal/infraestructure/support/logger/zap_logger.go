package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
    Logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
    // Configuración personalizada
    encoderConfig := zapcore.EncoderConfig{
    TimeKey:        "time",
    LevelKey:       "level",
    MessageKey:     "message",
    CallerKey:      "caller",
    StacktraceKey:  "stacktrace", // Mantener el stack trace para errores
    EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05"),
    EncodeLevel:    zapcore.CapitalLevelEncoder,
    EncodeCaller:   zapcore.ShortCallerEncoder,
    }

    // Configurar el nivel de log
    core := zapcore.NewCore(
    zapcore.NewConsoleEncoder(encoderConfig),
    zapcore.AddSync(zapcore.Lock(os.Stdout)),
    zap.LevelEnablerFunc(func(level zapcore.Level) bool {
        if level == zap.WarnLevel {
            encoderConfig.StacktraceKey = "" // Eliminar stack trace para Warn
        } else {
            encoderConfig.StacktraceKey = "stacktrace" // Mantener stack trace para Error
        }
        return true
    }),
    )

    logger := zap.New(core).Sugar()
    return &ZapLogger{Logger: logger}
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
