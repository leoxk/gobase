// File:        parser_test.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-04-18 by leoxiang

package xlsxpb

import (
	"testing"

	"github.com/leoxk/xlsxpb/test"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("Parse xlsx file", t, func() {
		p := NewParser()
		var config test.Config

		// with right field name
		So(p.Parse("./test/data.xlsx", "Table", "Tables", &config), ShouldBeNil)
		So(p.Parse("./test/data.xlsx", "AI", "AIs", &config), ShouldBeNil)
	})
}

// vim:ts=4:sw=4:et:ft=go:
