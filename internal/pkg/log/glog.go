package log

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	glog struct {
		logger *logrus.Entry
	}
)

func newGlog() *glog {
	return &glog{
		logger: newLogrusEntry(),
	}
}

// Info print info
func (g *glog) Info(args ...interface{}) {
	g.logger.Infoln(args...)
}

// Debugf print debug
func (g *glog) Debug(v ...interface{}) {
	g.logger.Debugln(v...)
}

// Warn print warning
func (g *glog) Warn(v ...interface{}) {
	g.logger.Warnln(v...)
}

// Errorf print error
func (g *glog) Error(v ...interface{}) {
	g.logger.Errorln(v...)
}

// Panic panic
func (g *glog) Panic(v ...interface{}) {
	g.logger.Panicln(v...)
}

// Infof print info with format.
func (g *glog) Infof(format string, v ...interface{}) {
	g.logger.Infof(format, v...)
}

// Debugf print debug with format.
func (g *glog) Debugf(format string, v ...interface{}) {
	g.logger.Debugf(format, v...)
}

// Warnf print warning with format.
func (g *glog) Warnf(format string, v ...interface{}) {
	g.logger.Warnf(format, v...)
}

// Errorf print error with format.
func (g *glog) Errorf(format string, v ...interface{}) {
	g.logger.Errorf(format, v...)
}

// Panicf panic with format.
func (g *glog) Panicf(format string, v ...interface{}) {
	g.logger.Panicf(format, v...)
}

func newLogrusEntry() *logrus.Entry {
	logger := logrus.New()
	logger.SetFormatter(formaterFromEnv())
	logger.SetLevel(levelFromEnv())
	logger.SetOutput(outputFromEnv())
	return logrus.NewEntry(logger)
}

func levelFromEnv() logrus.Level {
	lvl, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		lvl = logrus.DebugLevel
	}
	return lvl
}

func formaterFromEnv() logrus.Formatter {
	if strings.ToLower(os.Getenv("LOG_FORMAT")) == "json" {
		return &logrus.JSONFormatter{
			TimestampFormat: time.RFC1123,
		}
	}
	return &logrus.TextFormatter{
		TimestampFormat: time.RFC1123,
		FullTimestamp:   true,
	}
}

func outputFromEnv() io.WriteCloser {
	out := os.Getenv("LOG_OUTPUT")
	if strings.HasPrefix(out, filePrefix) {
		name := out[len(filePrefix):]
		f, err := os.Create(name)
		if err != nil {
			log.Printf("log: failed to create log file: %s, err: %v\n", name, err)
		}
		return f
	}
	return os.Stdout
}
