package log

var AppLog AppLogger

type AppLogger interface {
	Info(msg string)
	Error(msg string, err error)
	Debug(msg string)
	Warn(msg string)
}
