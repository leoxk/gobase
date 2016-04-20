// File:        to_slice.go
// Description: ---
// Notes:       ---
// Author:      leoxiang <leoxiang727@qq.com>
// Revision:    2016-04-20 by leoxiang

package base

import "reflect"

// ToSlice convert a typed slice to interface{} slice
// this will make sure you write easier generic code.
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

// vim:ts=4:sw=4:et:ft=go:
