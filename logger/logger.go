/*
 *	logger with json support
 *	Author: arthur
 *	Created Date: 2016-11-17
 *
 */

package logger

import ()

type Logger struct {
	name  string
	level Level
}

func GetLogger(name string) *Logger {
	return getLogger(name)
}

func (me *Logger) SetLevel(level Level) {
	me.level = level
}

func (me *Logger) SetLevelByName(level string) {
	me.SetLevel(getLevelByName(level))
}

func (me *Logger) Trace(format string, v ...interface{}) {
	if me.shouldLog(TRACE) {
		me.log(TRACE, format, v...)
	}
}

func (me *Logger) Debug(format string, v ...interface{}) {
	if me.shouldLog(DEBUG) {
		me.log(DEBUG, format, v...)
	}
}

func (me *Logger) Info(format string, v ...interface{}) {
	if me.shouldLog(INFO) {
		me.log(INFO, format, v...)
	}
}

func (me *Logger) Warn(format string, v ...interface{}) {
	if me.shouldLog(WARN) {
		me.log(WARN, format, v...)
	}
}

func (me *Logger) WarnError(err error) {
	if me.shouldLog(WARN) {
		me.log(WARN, "Error: %s", err.Error())
	}
}

func (me *Logger) Error(format string, v ...interface{}) {
	if me.shouldLog(ERROR) {
		me.log(ERROR, format, v...)
	}
}

func (me *Logger) ErrorError(err error) {
	if me.shouldLog(ERROR) {
		me.log(ERROR, "Error: %s", err.Error())
	}
}

func (me *Logger) Fatal(format string, v ...interface{}) {
	if me.shouldLog(FATAL) {
		me.log(FATAL, format, v...)
	}
}

func (me *Logger) FatalError(err error) {
	if me.shouldLog(FATAL) {
		me.log(FATAL, "Error: %s", err.Error())
	}
}

// Goa Logger

// logger used by Goa framework
var GoaLogger = GetLogger("goa")

// convenient functions

func Trace(format string, v ...interface{}) {
	GoaLogger.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	GoaLogger.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	GoaLogger.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	GoaLogger.Warn(format, v...)
}

func WarnError(err error) {
	GoaLogger.WarnError(err)
}

func Error(format string, v ...interface{}) {
	GoaLogger.Error(format, v...)
}

func ErrorError(err error) {
	GoaLogger.ErrorError(err)
}

func Fatal(format string, v ...interface{}) {
	GoaLogger.Fatal(format, v...)
}

func FatalError(err error) {
	GoaLogger.FatalError(err)
}

// global close

func SetRollingDaily(dir string, filename string) {
	logSetRollingDaily(dir, filename)
}

func Open(appRootPath string) {
	logOpen(appRootPath)
}

func Close() {
	logClose()
}
