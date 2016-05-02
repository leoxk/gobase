// File:        logger.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-02-03 by leoxiang

package log

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// A Level represents log output level.
// When logger output logs with lower level than predefined level,
// the log will be ignored.
type Level int

const (
	// LDebug messages are intended to help isolate a problem in a running system,
	// by showing the code that is executed, and the context information used during that execution.
	LDebug Level = iota

	// LInfo messages are intended to show what’s going on in the system, at a broad-brush level.
	LInfo

	// LWarn messages records that something in the system was not as expected.
	LWarn

	// LError messages records that something went wrong, i.e. some sort of failure occurred.
	LError

	// LFatal messages should generally only be used for recording a failure that prevents the system starting.
	LFatal

	// LMax is for array allocation
	LMax = LFatal
)

var l2s = [LMax + 1]string{"DBG ", "INF ", "WRN ", "ERR ", "FAT "}

// Logger define log instance.
type Logger struct {
	base   log.Logger
	bitmap [LMax + 1]bool
}

// NewLogger return new logger instance.
func NewLogger() *Logger {
	l := &Logger{}
	l.base.SetPrefix("")
	l.base.SetFlags(log.Lmicroseconds | log.Lshortfile)
	l.SetLevel(LDebug, LInfo, LWarn, LError, LFatal)
	l.SetWriter(os.Stderr)
	return l
}

// SetWriter set current log writer for log module
func (l *Logger) SetWriter(w io.Writer) {
	l.base.SetOutput(w)
}

// SetLevel set current log level
func (l *Logger) SetLevel(lvs ...Level) {
	for i := 0; i <= int(LMax); i++ {
		l.bitmap[i] = false
	}
	for _, lv := range lvs {
		l.bitmap[lv] = true
	}
}

// Debug calls Output to print to the standard logger.
func (l *Logger) Debug(format string, v ...interface{}) {
	l.output(LDebug, format, v...)
}

// Debugln calls Output to print to the standard logger.
func (l *Logger) Debugln(v ...interface{}) {
	l.outputln(LDebug, v...)
}

// Info calls Output to print to the standard logger.
func (l *Logger) Info(format string, v ...interface{}) {
	l.output(LInfo, format, v...)
}

// Warn calls Output to print to the standard logger.
func (l *Logger) Warn(format string, v ...interface{}) {
	l.output(LWarn, format, v...)
}

// Warnln calls Output to print to the standard logger.
func (l *Logger) Warnln(v ...interface{}) {
	l.outputln(LWarn, v...)
}

// Infoln calls Output to print to the standard logger.
func (l *Logger) Infoln(v ...interface{}) {
	l.outputln(LInfo, v...)
}

// Error calls Output to print to the standard logger.
func (l *Logger) Error(format string, v ...interface{}) {
	l.output(LError, format, v...)
}

// Errorln calls Output to print to the standard logger.
func (l *Logger) Errorln(v ...interface{}) {
	l.outputln(LError, v...)
}

// Fatal calls Output to print to the standard logger.
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.output(LFatal, format, v...)
}

// Fatalln calls Output to print to the standard logger.
func (l *Logger) Fatalln(v ...interface{}) {
	l.outputln(LFatal, v...)
}

// Panic calls Output to print to the standard logger.
func (l *Logger) Panic(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.base.SetPrefix(l2s[LFatal])
	l.base.Output(3, s)
	panic(s)
}

// Panicln calls Output to print to the standard logger.
func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.base.SetPrefix(l2s[LFatal])
	l.base.Output(3, s)
	panic(s)
}

// GenError calls Output to print to the standard logger.
func (l *Logger) GenError(format string, v ...interface{}) error {
	s := fmt.Sprintf(format, v...)
	l.base.SetPrefix(l2s[LError])
	l.base.Output(3, s)
	return errors.New(s)
}

// GenErrorln calls Output to print to the standard logger.
func (l *Logger) GenErrorln(v ...interface{}) error {
	s := fmt.Sprintln(v...)
	l.base.SetPrefix(l2s[LFatal])
	l.base.Output(3, s)
	return errors.New(s)
}

func (l *Logger) output(lv Level, format string, v ...interface{}) {
	if l.bitmap[lv] {
		l.base.SetPrefix(l2s[lv])
		l.base.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) outputln(lv Level, v ...interface{}) {
	if l.bitmap[lv] {
		l.base.SetPrefix(l2s[lv])
		l.base.Output(3, fmt.Sprintln(v...))
	}
}

// vim:ts=4:sw=4:et:ft=go:
