package log

var (
	root Logger
)

// Root return default logger instance
func Root() Logger {
	if root == nil {
		root = newGlog()
	}
	return root
}

// Info print info.
func Info(v ...interface{}) {
	Root().Info(v...)
}

// Debug print debug.
func Debug(v ...interface{}) {
	Root().Debug(v...)
}

// Warn print warning.
func Warn(v ...interface{}) {
	Root().Warn(v...)
}

// Error print error.
func Error(v ...interface{}) {
	Root().Error(v...)
}

// Panic panic.
func Panic(v ...interface{}) {
	Root().Panic(v...)
}

// Infof print info with format.
func Infof(format string, v ...interface{}) {
	Root().Infof(format, v...)
}

// Debugf print debug with format.
func Debugf(format string, v ...interface{}) {
	Root().Debugf(format, v...)
}

// Warnf print warning with format.
func Warnf(format string, v ...interface{}) {
	Root().Warnf(format, v...)
}

// Errorf print error with format.
func Errorf(format string, v ...interface{}) {
	Root().Errorf(format, v...)
}

// Panicf panic with format.
func Panicf(format string, v ...interface{}) {
	Root().Panicf(format, v...)
}
