## log

[![Build Status](https://travis-ci.org/leoxk/log.svg?branch=master)](https://travis-ci.org/leoxk/log)
[![GoDoc](https://godoc.org/github.com/leoxk/log?status.svg)](https://godoc.org/github.com/leoxk/log)
[![Go Report Card](http://goreportcard.com/badge/leoxk/log)](http://goreportcard.com/report/leoxk/log)

## Overview

a log wrapper for standard log, support auto rotate file mechanism.

## Usage

```
func TestLogger(t *testing.T) {
	l := NewLogger()

	Convey("Without setting, the default writer should be terminal", t, func() {
		l.SetLevel(LInfo)

		l.Info("test Info")
		l.Infoln("test Infoln")
		l.Error("test Error")
		l.Errorln("test Errorln")
		l.Fatal("test Fatal")
		l.Fatalln("test Fatalln")
	})

	Convey("After setting a file writer", t, func() {
		prefix := fmt.Sprintf("%s/test_rotate_file_writer.%d", os.TempDir(), rand.Int())
		w := NewRotateFileWriter(prefix)
		l.SetWriter(w)
	})
}
```
