/*
 *	logger with json support
 *	Author: arthur
 *	Created Date: 2016-11-17
 *
 */

package logger

import (
	"fmt"
	"sync"
	"time"
)

func (me *Logger) shouldLog(level Level) bool {
	return level >= me.level
}

type tLoggerMap struct {
	sync.RWMutex
	m map[string]*Logger
}

var loggerMap = tLoggerMap{m: make(map[string]*Logger)}

func getLogger(name string) *Logger {
	loggerMap.RLock()
	logger, ok := loggerMap.m[name]
	loggerMap.RUnlock()

	if !ok {
		loggerMap.Lock()
		logger = &Logger{name: name, level: defaultLogLevel}
		loggerMap.m[name] = logger
		loggerMap.Unlock()
	}

	return logger
}

// log functions

func getNow() string {
	t := time.Now()
	return t.Format("2006-01-02T15:04:05.000")
}

func (me *Logger) log(level Level, format string, v ...interface{}) {
	// TODO: write to file
	fmt.Printf("%s %s %s\n", getNow(), getLevelName(level), fmt.Sprintf(format, v...))
}

func logClose() {
	// TODO: close the log file
	Info("logger close")
}

// init the module

func init() {
	// TODO: open the log file
}
