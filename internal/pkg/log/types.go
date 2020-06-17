package log

import (
	"github.com/sirupsen/logrus"
)

type (
	Logger interface {
		// Infof print info with format.
		Infof(format string, v ...interface{})

		// Debugf print debug with format.
		Debugf(format string, v ...interface{})

		// Warnf print warning with format.
		Warnf(format string, v ...interface{})

		// Errorf print error with format.
		Errorf(format string, v ...interface{})

		// Panicf panic with format.
		Panicf(format string, v ...interface{})

		// Info print info.
		Info(v ...interface{})

		// Debug print debug.
		Debug(v ...interface{})

		// Warn print warning.
		Warn(v ...interface{})

		// Error print error.
		Error(v ...interface{})

		// Panic panic.
		Panic(v ...interface{})

		WithFields(fields Fields) Logger
	}

	glog struct {
		logger *logrus.Entry
	}

	Fields = map[string]interface{}

	contextKey string
)
