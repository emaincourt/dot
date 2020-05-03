package logger

type LogLevel = int

type LogLevelName = string

const DefaultLogLevel = "info"

const (
	LogLevelNameDebug LogLevelName = "debug"
	LogLevelNameInfo  LogLevelName = "info"
	LogLevelNameWarn  LogLevelName = "warn"
	LogLevelNameError LogLevelName = "error"
)

var LogLevelNameValueMap = map[string]int{
	LogLevelNameDebug: 1,
	LogLevelNameInfo:  2,
	LogLevelNameWarn:  3,
	LogLevelNameError: 4,
}

type Logger interface {
	Debug(message string)
	Debugf(message string, args ...interface{})
	Info(message string)
	Infof(message string, args ...interface{})
	Warn(message string)
	Warnf(message string, args ...interface{})
	Error(message string)
	Errorf(message string, args ...interface{})
	Table(header []string, data [][]string)
}
