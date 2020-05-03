package fmtlogger

import (
	"fmt"
	"os"

	"github.com/emaincourt/dot/pkg/logger"
	"github.com/olekukonko/tablewriter"
)

type FmtLogger struct {
	LogLevel logger.LogLevel
}

func NewFmtLogger(level string) *FmtLogger {
	return &FmtLogger{
		LogLevel: logger.LogLevelNameValueMap[level],
	}
}

func (l *FmtLogger) log(message string) {
	fmt.Println(message)
}

func (l *FmtLogger) shouldLog(level logger.LogLevelName) bool {
	return l.LogLevel <= logger.LogLevelNameValueMap[level]
}

func (l *FmtLogger) Debug(message string) {
	if l.shouldLog(logger.LogLevelNameDebug) {
		l.log(message)
	}
}

func (l *FmtLogger) Debugf(message string, args ...interface{}) {
	if l.shouldLog(logger.LogLevelNameDebug) {
		l.log(
			fmt.Sprintf(message, args...),
		)
	}
}

func (l *FmtLogger) Info(message string) {
	if l.shouldLog(logger.LogLevelNameInfo) {
		l.log(message)
	}
}

func (l *FmtLogger) Infof(message string, args ...interface{}) {
	if l.shouldLog(logger.LogLevelNameInfo) {
		l.log(
			fmt.Sprintf(message, args...),
		)
	}
}

func (l *FmtLogger) Warn(message string) {
	if l.shouldLog(logger.LogLevelNameWarn) {
		l.log(message)
	}
}

func (l *FmtLogger) Warnf(message string, args ...interface{}) {
	if l.shouldLog(logger.LogLevelNameWarn) {
		l.log(
			fmt.Sprintf(message, args...),
		)
	}
}

func (l *FmtLogger) Error(message string) {
	if l.shouldLog(logger.LogLevelNameError) {
		l.log(message)
	}
}

func (l *FmtLogger) Errorf(message string, args ...interface{}) {
	if l.shouldLog(logger.LogLevelNameError) {
		l.log(
			fmt.Sprintf(message, args...),
		)
	}
}

func (l *FmtLogger) Table(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
