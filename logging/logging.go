package logging

import (
	"log"
	"os"
)

// Type Flags represents logger flags (see log.L* flags).
type Flags int8

var (
	// More verbose logging flags for debugging/development.
	FlagsDevelopment Flags = log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix
	// Less verbose logging info.
	FlagsProduction Flags = log.Ldate | log.Ltime | log.Lmsgprefix

	loggingLevel = InfoLevel
	logger       = log.New(os.Stderr, "", int(FlagsProduction))
)

// Set the minimum logging level for the default logger.
func SetLevel(lvl Level) {
	loggingLevel = lvl
}

// Change the flags of the default logger. The default flags are FlagsProduction.
func SetFlags(flags Flags) {
	logger.SetFlags(int(flags))
}

// Checks level and does the logging.
func do(lvl Level, caller string, template string, args ...interface{}) {
	if lvl < loggingLevel {
		return
	}

	// Gotta check if the flags we're using need the prefix or not, to save up on locking the logger
	if caller != "" && logger.Flags()&log.Lmsgprefix > 0 {
		logger.SetPrefix("[" + caller + "] ")
		defer logger.SetPrefix("")
	}

	if template == "" {
		logger.Print(args...)
	} else {
		logger.Printf(template, args...)
	}

	// This is actually just as fast as using log.Panic and log.Fatal
	switch lvl {
	case PanicLevel:
		panic(caller)
	case FatalLevel:
		os.Exit(1)
	}
}

// Debug uses fmt.Sprint to construct and log a message at DebugLevel.
func Debug(caller string, args ...interface{}) {
	do(DebugLevel, caller, "", args...)
}

// Info uses fmt.Sprint to construct and log a message at InfoLevel.
func Info(caller string, args ...interface{}) {
	do(InfoLevel, caller, "", args...)
}

// Warn uses fmt.Sprint to construct and log a message at WarnLevel.
func Warn(caller string, args ...interface{}) {
	do(WarnLevel, caller, "", args...)
}

// Error uses fmt.Sprint to construct and log a message at ErrorLevel.
func Error(caller string, args ...interface{}) {
	do(ErrorLevel, caller, "", args...)
}

// Panic uses fmt.Sprint to construct and log a message at PanicLevel, then panics.
func Panic(caller string, args ...interface{}) {
	do(PanicLevel, caller, "", args...)
}

// Fatal uses fmt.Sprint to construct and log a message at FatalLevel, then calls os.Exit.
func Fatal(caller string, args ...interface{}) {
	do(FatalLevel, caller, "", args...)
}

// Debugf uses fmt.Sprintf to log a formatted message at DebugLevel.
func Debugf(caller string, template string, args ...interface{}) {
	do(DebugLevel, caller, template, args...)
}

// Infof uses fmt.Sprintf log a formatted message at InfoLevel.
func Infof(caller string, template string, args ...interface{}) {
	do(InfoLevel, caller, template, args...)
}

// Warnf uses fmt.Sprintf log a formatted message at WarnLevel.
func Warnf(caller string, template string, args ...interface{}) {
	do(WarnLevel, caller, template, args...)
}

// Errorf uses fmt.Sprintf log a formatted message at ErrorLevel.
func Errorf(caller string, template string, args ...interface{}) {
	do(ErrorLevel, caller, template, args...)
}

// Panicf uses fmt.Sprintf log a formatted message at PanicLevel, then panics.
func Panicf(caller string, template string, args ...interface{}) {
	do(PanicLevel, caller, template, args...)
}

// Fatalf uses fmt.Sprintf log a formatted message at FatalLevel, then calls os.Exit.
func Fatalf(caller string, template string, args ...interface{}) {
	do(FatalLevel, caller, template, args...)
}
