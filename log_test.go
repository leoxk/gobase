// File:        log_test.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-02-03 by leoxiang

package log

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
		// l.SetLevel()

		l.Debug("test Debug")
		l.Debugln("test Debugln")
		l.Info("test Info")
		l.Infoln("test Infoln")
		l.Warn("test Warn")
		l.Warnln("test Warnln")
		l.Error("test Error")
		l.Errorln("test Errorln")
		l.Fatal("test Fatal")
		l.Fatalln("test Fatalln")

		So(l.GenError("test return error").Error(), ShouldEqual, "test return error")
	})

	Convey("After setting a file writer", t, func() {
		prefix := fmt.Sprintf("%s/test_rotate_file_writer.%d", os.TempDir(), rand.Int())
		w := NewRotateFileWriter(prefix)
		l.SetWriter(w)
	})
}

// vim:ts=4:sw=4:et:ft=go:
