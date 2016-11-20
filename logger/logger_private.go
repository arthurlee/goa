/*
 *	logger with json support
 *	Author: arthur
 *	Created Date: 2016-11-17
 *
 */

package logger

import (
	"fmt"
	"os"
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

func getToday() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func getTodayInt() int {
	year, month, day := time.Now().Date()
	return year*10000 + int(month)*100 + day
}

// init the module

type logFile struct {
	appRootPath string
	dir         string
	filename    string
	today       int
	file        *os.File
	ch          chan string
	console     bool
	stop_ch     chan bool
}

var log_file logFile

func (me *Logger) log(level Level, format string, v ...interface{}) {
	message := fmt.Sprintf("%s %s [%s] %s\n", getNow(), getLevelName(level), me.name, fmt.Sprintf(format, v...))
	log_file.ch <- message
}

func logSetRollingDaily(dir string, filename string) {
	log_file.dir = dir
	log_file.filename = filename
}

func logOpen(appRootPath string) {
	log_file.appRootPath = appRootPath
	log_file.ch = make(chan string, 4096)
	log_file.stop_ch = make(chan bool, 1)
	log_file.console = true

	go logHandler(log_file.ch)
}

func closeFile() {
	if log_file.file != nil {
		err := log_file.file.Sync()
		if err != nil {
			fmt.Println(err)
		}
		err = log_file.file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func openNewFile() error {
	filename := fmt.Sprintf("%s/%s/%s.%s", log_file.appRootPath, log_file.dir, log_file.filename, getToday())
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	log_file.file = file
	return err
}

func logWrite(message string) {
	today := getTodayInt()
	if log_file.today != today {
		closeFile()

		err := openNewFile()
		if err != nil {
			return
		}

		log_file.today = today
	}

	_, err := log_file.file.WriteString(message)
	if err != nil {
		fmt.Println(err)
	}

	if log_file.console {
		fmt.Printf(message)
	}
}

func logClose() {
	close(log_file.ch)

	// wait log handler exit
	_, ok := <-log_file.stop_ch
	if !ok {
		fmt.Println("fail to wait log handler exit")
	}
}

func logHandler(ch chan string) {
	logWrite(getNow() + " INFO [Goa Logger] --------------- start ---------------\n")

	for {
		message, ok := <-ch
		if !ok {
			break
		}

		logWrite(message)
	}

	logWrite(getNow() + " INFO [Goa Logger] --------------- end ---------------\n")
	closeFile()

	log_file.stop_ch <- true
}
