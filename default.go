// File:        default.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-03-01 by leoxiang

package log

import (
	"errors"
	"fmt"
	"io"
)

var l = NewLogger()

// SetWriter set current log writer for log module
func SetWriter(w io.Writer) {
	l.base.SetOutput(w)
}

// SetLevel set current log level
func SetLevel(lvs ...Level) {
	l.SetLevel(lvs...)
}

// Debug calls Output to print to the standard logger.
func Debug(format string, v ...interface{}) {
	l.output(LDebug, format, v...)
}

// Debugln calls Output to print to the standard logger.
func Debugln(v ...interface{}) {
	l.outputln(LDebug, v...)
}

// Info calls Output to print to the standard logger.
func Info(format string, v ...interface{}) {
	l.output(LInfo, format, v...)
}

// Infoln calls Output to print to the standard logger.
func Infoln(v ...interface{}) {
	l.outputln(LInfo, v...)
}

// Warn calls Output to print to the standard logger.
func Warn(format string, v ...interface{}) {
	l.output(LWarn, format, v...)
}

// Warnln calls Output to print to the standard logger.
func Warnln(v ...interface{}) {
	l.outputln(LWarn, v...)
}

// Error calls Output to print to the standard logger.
func Error(format string, v ...interface{}) {
	l.output(LError, format, v...)
}

// Errorln calls Output to print to the standard logger.
func Errorln(v ...interface{}) {
	l.outputln(LError, v...)
}

// Fatal calls Output to print to the standard logger.
func Fatal(format string, v ...interface{}) {
	l.output(LFatal, format, v...)
}

// Fatalln calls Output to print to the standard logger.
func Fatalln(v ...interface{}) {
	l.outputln(LFatal, v...)
}

// Panic calls Output to print to the standard logger.
func Panic(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.base.SetPrefix(l2s[LFatal])
	l.base.Output(3, s)
	panic(s)
}

// Panicln calls Output to print to the standard logger.
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.base.SetPrefix(l2s[LFatal])
	l.base.Output(3, s)
	panic(s)
}

// GenError calls Output to print to the standard logger.
func GenError(format string, v ...interface{}) error {
	s := fmt.Sprintf(format, v...)
	l.base.SetPrefix(l2s[LError])
	l.base.Output(3, s)
	return errors.New(s)
}

// GenErrorln calls Output to print to the standard logger.
func GenErrorln(v ...interface{}) error {
	s := fmt.Sprintln(v...)
	l.base.SetPrefix(l2s[LFatal])
	l.base.Output(3, s)
	return errors.New(s)
}

// vim:ts=4:sw=4:et:ft=go:
