package logger

import (
	"strings"
)

type Level int

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

var defaultLogLevel Level = DEBUG

func getLevelByName(levelName string) Level {
	levelName = strings.ToUpper(levelName)

	level := defaultLogLevel
	for i, v := range levelNames {
		if v == levelName {
			level = Level(i)
			break
		}
	}

	return level
}

func getLevelName(level Level) string {
	if level >= TRACE && level <= FATAL {
		return levelNames[level]
	}
	return "Unkown"
}
