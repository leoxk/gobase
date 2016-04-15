// File:        logger.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-02-03 by leoxiang

package log

import (
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
	LInfo Level = iota
	LError
	LFatal
	lMax = LFatal
)

var l2s = [lMax + 1]string{"INF ", "ERR ", "FAT "}

// Logger define log instance.
type Logger struct {
	base  log.Logger
	level Level
}

// NewLogger return new logger instance.
func NewLogger() *Logger {
	l := &Logger{}
	l.base.SetPrefix("")
	l.base.SetFlags(log.Lmicroseconds | log.Lshortfile)
	l.SetLevel(LInfo)
	l.SetWriter(os.Stderr)
	return l
}

// SetWriter set current log writer for log module
func (l *Logger) SetWriter(w io.Writer) {
	l.base.SetOutput(w)
}

// SetLevel set current log level
func (l *Logger) SetLevel(lv Level) {
	l.level = lv
}

// Info calls Output to print to the standard logger.
func (l *Logger) Info(format string, v ...interface{}) {
	l.output(LInfo, format, v...)
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

func (l *Logger) output(lv Level, format string, v ...interface{}) {
	if lv >= l.level {
		l.base.SetPrefix(l2s[lv])
		l.base.Output(3, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) outputln(lv Level, v ...interface{}) {
	if lv >= l.level {
		l.base.SetPrefix(l2s[lv])
		l.base.Output(3, fmt.Sprintln(v...))
	}
}

// vim:ts=4:sw=4:et:ft=go:
