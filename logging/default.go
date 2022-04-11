package logging

// The default/global logger instance.
var logger = New()

// Debug uses fmt.Sprint to construct and log a message at DebugLevel.
func Debug(args ...interface{}) {
	logger.do(DebugLevel, "", args...)
}

// Info uses fmt.Sprint to construct and log a message at InfoLevel.
func Info(args ...interface{}) {
	logger.do(InfoLevel, "", args...)
}

// Warn uses fmt.Sprint to construct and log a message at WarnLevel.
func Warn(args ...interface{}) {
	logger.do(WarnLevel, "", args...)
}

// Error uses fmt.Sprint to construct and log a message at ErrorLevel.
func Error(args ...interface{}) {
	logger.do(ErrorLevel, "", args...)
}

// Panic uses fmt.Sprint to construct and log a message at PanicLevel, then panics.
func Panic(args ...interface{}) {
	logger.do(PanicLevel, "", args...)
}

// Fatal uses fmt.Sprint to construct and log a message at FatalLevel, then calls os.Exit.
func Fatal(args ...interface{}) {
	logger.do(FatalLevel, "", args...)
}

// Debugf uses fmt.Sprintf to log a formatted message at DebugLevel.
func Debugf(template string, args ...interface{}) {
	logger.do(DebugLevel, template, args...)
}

// Infof uses fmt.Sprintf log a formatted message at InfoLevel.
func Infof(template string, args ...interface{}) {
	logger.do(InfoLevel, template, args...)
}

// Warnf uses fmt.Sprintf log a formatted message at WarnLevel.
func Warnf(template string, args ...interface{}) {
	logger.do(WarnLevel, template, args...)
}

// Errorf uses fmt.Sprintf log a formatted message at ErrorLevel.
func Errorf(template string, args ...interface{}) {
	logger.do(ErrorLevel, template, args...)
}

// Panicf uses fmt.Sprintf log a formatted message at PanicLevel, then panics.
func Panicf(template string, args ...interface{}) {
	logger.do(PanicLevel, template, args...)
}

// Fatalf uses fmt.Sprintf log a formatted message at FatalLevel, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	logger.do(FatalLevel, template, args...)
}
