# Go Base
[![Build Status](https://travis-ci.org/leoxk/gobase.svg?branch=master)](https://travis-ci.org/leoxk/gobase)
[![GoDoc](https://godoc.org/github.com/leoxk/gobase?status.svg)](https://godoc.org/github.com/leoxk/gobase)

# Overview
--------
### Log
Simple log wrapper for standard log, support auto rotate file mechanism.

### xlsxpb
Parse xlsx into protobuf message.

#Usage
```
func TestRotateFileWriter(t *testing.T) {
	Convey("Open new rotate file writer in system tmp dir", t, func() {
		prefix := fmt.Sprintf("%s/test_rotate_file_writer.%d", os.TempDir(), rand.Int())
		w := NewRotateFileWriter(prefix)

		Convey("The writer should not be nil", func() {
			So(w, ShouldNotBeNil)
		})

		Convey("The writer current file state should be regular", func() {
			info, err := os.Stat(w.file.Name())
			So(err, ShouldBeNil)
			So(info.Mode().IsRegular(), ShouldBeTrue)
		})
	})
}

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
