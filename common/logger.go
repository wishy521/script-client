package common

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var Log = NewLogger()

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Debug(method string, v ...interface{}) {
	l.print("INFO", method, v...)
}

func (l *Logger) Info(method string, v ...interface{}) {
	l.print("INFO", method, v...)
}

func (l *Logger) Error(method string, v ...interface{}) {
	l.print("ERROR", method, v...)
}


func (l *Logger) print(level string, method string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	fileInfo := fmt.Sprintf("%s:%d", file, line)
	l.logger.SetPrefix(fmt.Sprintf("[%s] %s ", level, method))
	l.logger.Output(3, fmt.Sprintf("%s %s", fileInfo, fmt.Sprint(v...)))
}
