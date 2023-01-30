//loggerinterface.go - содержит описание интерфейса которому должно соответствовать
//применняемое средство логирования

package logger

//Logger - интерфейс которому должно соответствовать средство логирования
type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}
