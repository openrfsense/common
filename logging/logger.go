// Copyright (C) 2022 OpenRFSense
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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

// Type Logger represents a logger instance with a specific unit name (prefix) and
// logging level. New instances are to be created with logging.New().
type Logger struct {
	logger *log.Logger
	lvl    Level
	name   string
}

// Creates a new Logger with the given options. The default logger has FlagsProduction,
// logs at InfoLevel, has no prefix and outputs to os.Stderr.
func New() *Logger {
	ret := &Logger{
		logger: log.New(os.Stderr, "", int(FlagsProduction)),
		lvl:    InfoLevel,
		name:   "",
	}

	return ret
}

// Changes the flags for the inderlying log.Logger. Default is FlagsProduction.
func (l *Logger) WithFlags(flags Flags) *Logger {
	l.logger.SetFlags(int(flags))
	return l
}

// Sets a minimum logging level for the logger being created. Default is InfoLevel.
func (l *Logger) WithLevel(lvl Level) *Logger {
	l.lvl = lvl
	return l
}

// Sets the output sink for the underlying log.Logger.
func (l *Logger) WithOutput(w io.Writer) *Logger {
	l.logger.SetOutput(w)
	return l
}

// Gives a specific name to the logger. Will be included in the output as '[prefix]'.
// Default is empty string (no prefix will be printed).
func (l *Logger) WithPrefix(prefix string) *Logger {
	l.name = prefix
	l.logger.SetPrefix(fmt.Sprintf("[%s] ", prefix))
	return l
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
