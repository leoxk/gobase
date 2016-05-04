// File:        writer.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-02-03 by leoxiang

package log

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// RotateFileWriter define log write with level module and file rotate mechanism.
type RotateFileWriter struct {
	prefix string
	file   *os.File
	ctime  time.Time
}

// NewRotateFileWriter create and return a new RotateFileWriter.
func NewRotateFileWriter(p string) *RotateFileWriter {
	// init data
	var tmp RotateFileWriter
	tmp.prefix, _ = filepath.Abs(p)

	// create the first file
	tmp.create()

	return &tmp
}

func (w *RotateFileWriter) rotate() {
	// for new instance just do rotate
	if w.file == nil {
		w.create()
		return
	}

	// compare with now
	if w.ctime.Year() == time.Now().Year() ||
		w.ctime.Month() == time.Now().Month() ||
		w.ctime.Day() == time.Now().Day() {
		return
	}

	// shift old
	oldPath := w.file.Name()
	newPath := fmt.Sprintf("%s/%d/%d/%s",
		filepath.Base(oldPath),
		w.ctime.Year(),
		w.ctime.Month(),
		filepath.Dir(oldPath))
	if err := os.Rename(oldPath, newPath); err != nil {
		return
	}

	// create
	w.create()
}

func (w *RotateFileWriter) create() {
	// release old
	if w.file != nil {
		w.file.Close()
	}
	w.file = nil

	// open new file
	path := w.prefix + "." + time.Now().Format("2006-01-02")
	flag := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	perm := os.FileMode(0644)
	if tmp, err := os.OpenFile(path, flag, perm); err == nil {
		w.file = tmp
		w.ctime = time.Now()
	}
}

func (w *RotateFileWriter) Write(p []byte) (n int, err error) {
	// check rotate
	w.rotate()

	// check nil
	if w.file == nil {
		return 0, errors.New("logger: file is nil")
	}

	return w.file.Write(p)
}

// vim:ts=4:sw=4:et:ft=go:
