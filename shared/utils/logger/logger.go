package logger

import "log"

func Debug(msg string, args ...interface{}) {
	logArgs := []interface{}{msg}
	logArgs = append(logArgs, args...)
	log.Print(logArgs...)
}
