package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Type Flags represents logger flags (a wrapper around log.L* flags).
type Flags int8

var (
	// More verbose logging flags for debugging/development.
	FlagsDevelopment Flags = log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix
	// Less verbose logging info.
	FlagsProduction Flags = log.LstdFlags | log.Lmsgprefix
)

// Type Option is a function which takes a Logger instance and modifies some of its
// internal configuration.
type Option func(*Logger)

// Type Logger represents a logger instance with a specific unit name (prefix) and
// logging level. New instances are to be created with logging.New().
type Logger struct {
	logger *log.Logger
	lvl    Level
	name   string
}

// Creates a new Logger with the given options. The default logger has FlagsProduction,
// logs at InfoLevel, has no prefix and outputs to os.Stderr.
func New(options ...Option) *Logger {
	ret := &Logger{
		logger: log.New(os.Stderr, "", int(FlagsProduction)),
		lvl:    InfoLevel,
		name:   "",
	}

	for _, opt := range options {
		opt(ret)
	}

	return ret
}

// Gives a specific name to the logger. Will be included in the output as '[prefix]'.
// Default is empty string (no prefix will be printed).
func WithPrefix(prefix string) Option {
	return func(l *Logger) {
		l.name = prefix
		l.logger.SetPrefix(fmt.Sprintf("[%s] ", prefix))
	}
}

// Sets a minimum logging level for the logger being created. Default is InfoLevel.
func WithLevel(lvl Level) Option {
	return func(l *Logger) {
		l.lvl = lvl
	}
}

// Changes the flags for the inderlying log.Logger. Default is FlagsProduction.
func WithFlags(flags Flags) Option {
	return func(l *Logger) {
		l.logger.SetFlags(int(flags))
	}
}

// Sets the output sink for the underlying log.Logger.
func WithOutput(w io.Writer) Option {
	return func(l *Logger) {
		l.logger.SetOutput(w)
	}
}

// Does the logging and panic/Exit.
func (l Logger) do(lvl Level, template string, args ...interface{}) {
	if lvl < l.lvl {
		return
	}

	if template == "" {
		l.logger.Print(args...)
	} else {
		l.logger.Printf(template, args...)
	}

	// This is actually just as fast as using log.Panic and log.Fatal
	switch lvl {
	case PanicLevel:
		panic(l.name)
	case FatalLevel:
		os.Exit(1)
	}
}

// Debug uses fmt.Sprint to construct and log a message at DebugLevel.
func (l Logger) Debug(args ...interface{}) {
	l.do(DebugLevel, "", args...)
}

// Info uses fmt.Sprint to construct and log a message at InfoLevel.
func (l Logger) Info(args ...interface{}) {
	l.do(InfoLevel, "", args...)
}

// Warn uses fmt.Sprint to construct and log a message at WarnLevel.
func (l Logger) Warn(args ...interface{}) {
	l.do(WarnLevel, "", args...)
}

// Error uses fmt.Sprint to construct and log a message at ErrorLevel.
func (l Logger) Error(args ...interface{}) {
	l.do(ErrorLevel, "", args...)
}

// Panic uses fmt.Sprint to construct and log a message at PanicLevel, then panics.
func (l Logger) Panic(args ...interface{}) {
	l.do(PanicLevel, "", args...)
}

// Fatal uses fmt.Sprint to construct and log a message at FatalLevel, then calls os.Exit.
func (l Logger) Fatal(args ...interface{}) {
	l.do(FatalLevel, "", args...)
}

// Debugf uses fmt.Sprintf to log a formatted message at DebugLevel.
func (l Logger) Debugf(template string, args ...interface{}) {
	l.do(DebugLevel, template, args...)
}

// Infof uses fmt.Sprintf log a formatted message at InfoLevel.
func (l Logger) Infof(template string, args ...interface{}) {
	l.do(InfoLevel, template, args...)
}

// Warnf uses fmt.Sprintf log a formatted message at WarnLevel.
func (l Logger) Warnf(template string, args ...interface{}) {
	l.do(WarnLevel, template, args...)
}

// Errorf uses fmt.Sprintf log a formatted message at ErrorLevel.
func (l Logger) Errorf(template string, args ...interface{}) {
	l.do(ErrorLevel, template, args...)
}

// Panicf uses fmt.Sprintf log a formatted message at PanicLevel, then panics.
func (l Logger) Panicf(template string, args ...interface{}) {
	l.do(PanicLevel, template, args...)
}

// Fatalf uses fmt.Sprintf log a formatted message at FatalLevel, then calls os.Exit.
func (l Logger) Fatalf(template string, args ...interface{}) {
	l.do(FatalLevel, template, args...)
}
