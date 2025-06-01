package logger

type Logger interface {
	Info(args ...any)
	Error(args ...any)
	Debug(args ...any)
	Warn(args ...any)
	Infow(msg string, keysAndValues ...any)
	Errorw(msg string, keysAndValues ...any)
	Warnf(msg string, keysAndValues ...any)
}
