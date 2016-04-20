// File:        parser.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-04-18 by leoxiang

// Package xlsxpb is a xlsx parser which convert xlsx into a protobuf message
package xlsxpb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/leoxk/gobase/log"
	"github.com/tealeg/xlsx"
)

// Parser define a parser instance.
type Parser struct {
	// StartLine should be the line to start parsing,
	// the first line should contain the meta information
	// and the others below should consist the data.
	StartLine int

	fileName   string
	sheetName  string
	rowNumber  int
	cellNumber int
}

const (
	// DefDelimitToken is default delimit name used to
	// split different token in a full field name.
	DefDelimitToken string = "."
)

// NewParser create and return a new parser.
func NewParser() *Parser {
	return &Parser{
		StartLine: 0,
	}
}

func (p *Parser) String() string {
	return fmt.Sprintf("[%s:%s:%d:%d]", p.fileName, p.sheetName, p.rowNumber, p.cellNumber)
}

// Parse read and convert a xlsx file into protobuf message.
func (p *Parser) Parse(fileName, sheetName, fieldName string, message proto.Message) error {
	// open file
	file, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Errorln("open input file failed", fileName)
		return err
	}
	p.fileName = fileName

	// get sheet
	var sheet *xlsx.Sheet
	if sheetName == "" {
		// empty then use the first sheet
		sheet = file.Sheets[0]
	} else {
		// try find the sheet
		var ok bool
		if sheet, ok = file.Sheet[sheetName]; !ok {
			e := fmt.Errorf("xlsx do not contain sheet %s %s", fileName, sheetName)
			log.Errorln(e.Error())
			return e
		}
	}
	p.sheetName = sheetName

	// get meta line
	if p.StartLine >= len(sheet.Rows) {
		e := fmt.Errorf("invalid start line %d %d", p.StartLine, len(sheet.Rows))
		log.Errorln(e.Error())
		return e
	}
	metaRow := sheet.Rows[p.StartLine]

	// iterate all row
	for i := p.StartLine + 1; i < sheet.MaxRow; i++ {
		p.rowNumber = i
		if err = p.parseRow(metaRow, sheet.Rows[i], fieldName, message); err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseRow(meta, data *xlsx.Row, fieldName string, message proto.Message) error {
	// if message is not struct
	v := reflect.ValueOf(message).Elem()
	if v.Kind() != reflect.Struct {
		e := fmt.Errorf("message is not struct %v %v", p, v.Kind())
		log.Errorln(e.Error())
		return e
	}

	// get field from name
	f := v.FieldByName(fieldName)
	if !f.IsValid() {
		e := fmt.Errorf("cant find field in message %v %v %s", p, v.Type(), fieldName)
		log.Errorln(e.Error())
		return e
	}

	// iterate all cells
	for i, cell := range data.Cells {
		p.cellNumber = i
		// ignore empty
		if cell.Value == "" {
			continue
		}

		// check field name
		if i >= len(meta.Cells) || meta.Cells[i].Value == "" {
			e := fmt.Errorf("cell dont have fieldname %v", p)
			log.Errorln(e.Error())
			return e
		}
		// fieldName := meta.Cells[i].Value

		// if err := p.parseCell(fieldName, cell.Value, f); err != nil {
		// return err
		// }
	}

	return nil
}

func (p *Parser) parseCell(fullFieldName, dataValue string, message proto.Message) error {
	// iterate all tokens
	msg := message
	for _, tok := range strings.Split(fullFieldName, DefDelimitToken) {
		// get reflect value
		v := reflect.ValueOf(msg).Elem()

		// get field type
		f := v.FieldByName(tok)
		if !f.IsValid() {
			e := fmt.Errorf("invalid fieldname %v %v %s", p, v.Type(), tok)
			log.Errorln(e.Error())
			return e
		}
	}

	return nil
}

// vim:ts=4:sw=4:et:ft=go:
