package log

var AppLogger interface {
	Info(msg string)
	Error(msg string, err error)
	Debug(msg string)
	Warn(msg string)
}
