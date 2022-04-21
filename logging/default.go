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
