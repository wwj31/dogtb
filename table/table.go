package table

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

type Table struct {
	columnsName []string
	rows        [][]string
	style       int
}

func Create(v ...interface{}) (*Table, error) {
	tb := &Table{}
	if len(v) == 0 {
		return tb, nil
	}

	v0 := v[0]
	kind := reflect.TypeOf(v0).Kind()
	switch kind {
	case reflect.String:
		if !checkAllInOne(v) {
			return nil, fmt.Errorf("not all-in-one")
		}
		for _, colName := range v {
			tb.addColumn(colName.(string))
		}
	case reflect.Struct, reflect.Pointer:
		if !checkAllInOne(v) {
			return nil, fmt.Errorf("not all-in-one")
		}

		var v0typ reflect.Type
		if kind == reflect.Pointer {
			v0typ = reflect.ValueOf(v0).Elem().Type()
		} else {
			v0typ = reflect.TypeOf(v0)
		}

		for i := 0; i < v0typ.NumField(); i++ {
			structFiled := v0typ.Field(i)
			name := structFiled.Name
			if tag, ok := structFiled.Tag.Lookup("dogtg"); ok {
				name = tag
			}
			tb.columnsName = append(tb.columnsName, name)
		}

		var (
			typ reflect.Type
			val reflect.Value
		)

		for _, elem := range v {
			if kind == reflect.Pointer {
				typ = reflect.ValueOf(elem).Elem().Type()
				val = reflect.ValueOf(elem).Elem()
			} else {
				typ = reflect.TypeOf(elem)
				val = reflect.ValueOf(elem)
			}

			if typ.String() != v0typ.String() {
				return nil, fmt.Errorf("struct type not all-in-one")
			}

			var row []string
			for i := 0; i < typ.NumField(); i++ {
				valueField := val.Field(i)
				if valueField.CanInterface() {
					row = append(row, cast.ToString(valueField.Interface()))
				}
			}
			tb.addRow(row)
		}

	default:
		return nil, fmt.Errorf("invalid king %v", kind)
	}
	return tb, nil
}

/*
┌┬┐  ┌─┐
├┼┤  │┼│
└┴┘  └─┘
*/
func (t *Table) String() string {
	return t.line1()
}

func (t *Table) addColumn(name string) {
	t.columnsName = append(t.columnsName, name)
}

func (t *Table) addRow(row []string) {
	t.rows = append(t.rows, row)
}

func (t *Table) line1() string {
	str := "┌"
	for n, name := range t.columnsName {
		for i := 0; i < stringWidth(name); i++ {
			str += "─"
		}

		if n == len(t.columnsName)-1 {
			str += "┐"
		} else {
			str += "┬"
		}
	}
	return str
}

func stringWidth(str string) int {
	return len(str) + 2
}
