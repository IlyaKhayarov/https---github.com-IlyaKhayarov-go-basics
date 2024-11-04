package main

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtend struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogExtended() *LogExtend {
	return &LogExtend{Logger: log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError}
}

func (l *LogExtend) Infoln(msg string) {
	l.println(LogLevelInfo, "INFO ", msg)
}

func (l *LogExtend) Warnln(msg string) {
	l.println(LogLevelWarning, "WARN ", msg)
}

func (l *LogExtend) Errorln(msg string) {
	l.println(LogLevelError, "ERR ", msg)
}

func (l *LogExtend) SetLogLevel(logLvl LogLevel) {
	l.logLevel = logLvl
}

func (l *LogExtend) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}

func main() {
	logger := NewLogExtended()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}
