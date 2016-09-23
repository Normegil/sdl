package log

import "github.com/Sirupsen/logrus"

// StructuredLog support decorate Logrus to implement AgnosticLogger
type StructuredLog struct {
	Logger logrus.FieldLogger
}

// Log log a message to the output defined in logrus.
func (l StructuredLog) Log(lvl Level, v ...interface{}) {
	if nil != l.Logger {
		switch lvl {
		case PANIC:
			l.Logger.Panic(v...)
		case FATAL:
			l.Logger.Fatal(v...)
		case ERROR:
			l.Logger.Error(v...)
		case WARN:
			l.Logger.Warn(v...)
		case INFO:
			l.Logger.Info(v...)
		case DEBUG:
			l.Logger.Debug(v...)
		case TRACE:
			l.Logger.WithField("level", "trace").Print(v...)
		}
	}
}

// With send back a logger containing the fields in the given structure
func (l StructuredLog) With(str Structure) AgnosticLogger {
	fields := logrus.Fields{}
	for key, value := range str {
		fields[key] = value
	}

	l.Logger = l.Logger.WithFields(fields)
	return l
}
