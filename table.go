package dogtb

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

const width = 2

type Table struct {
	columnsName   []string
	columnsMaxLen []int
	rows          [][]string
	style         int
}

func Create(v ...interface{}) (*Table, error) {
	tb := &Table{}
	if len(v) == 0 {
		return tb, nil
	}

	v0 := v[0]
	kind := reflect.TypeOf(v0).Kind()
	switch kind {
	case reflect.Slice:
		v0Value := reflect.ValueOf(v0)
		v0Len := v0Value.Len()
		slice := reflect.MakeSlice(reflect.TypeOf(make([]interface{}, 1)), v0Len, v0Len)
		for i := 0; i < v0Len; i++ {
			slice.Index(i).Set(v0Value.Index(i))
		}
		arr, _ := slice.Interface().([]interface{})
		return Create(arr...)
	case reflect.String:
		if !checkAllInOne(v) {
			return nil, fmt.Errorf("not all-in-one")
		}
		for _, colName := range v {
			tb.addColumn(colName.(string))
		}
	case reflect.Struct, reflect.Ptr:
		if !checkAllInOne(v) {
			return nil, fmt.Errorf("not all-in-one")
		}

		var v0typ reflect.Type
		if kind == reflect.Ptr {
			v0typ = reflect.ValueOf(v0).Elem().Type()
		} else {
			v0typ = reflect.TypeOf(v0)
		}

		for i := 0; i < v0typ.NumField(); i++ {
			structFiled := v0typ.Field(i)
			name := structFiled.Name
			if tag, ok := structFiled.Tag.Lookup("tb"); ok {
				name = tag
			}
			tb.columnsName = append(tb.columnsName, name)
		}

		var (
			typ reflect.Type
			val reflect.Value
		)

		for _, elem := range v {
			if kind == reflect.Ptr {
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

// String get table format
func (t *Table) String(style ...int) string {
	if len(style) != 0 {
		t.style = style[0]
	}
	t.maxLen()

	result := t.line(0) + "\n"
	result += t.rowData(t.columnsName) + "\n"
	result += t.line(1) + "\n"

	for _, row := range t.rows {
		result += t.rowData(row) + "\n"
		//result += t.line(1) + "\n"
	}

	result += t.line(2) + "\n"
	return result
}

func (t *Table) addColumn(name string) {
	t.columnsName = append(t.columnsName, name)
}

func (t *Table) addRow(row []string) {
	t.rows = append(t.rows, row)
}
func (t *Table) maxLen() {
	stats := append(t.rows, t.columnsName)

	var maxLen []int
	for i := 0; i < len(t.columnsName); i++ {
		var max int
		for j := 0; j < len(stats); j++ {
			if len(stats[j][i]) > max {
				max = len(stats[j][i])
			}
		}

		maxLen = append(maxLen, max)
	}

	t.columnsMaxLen = maxLen
}

// line type valid in(0,1,2),define top and middle and bottom of frame
func (t *Table) line(typ int) string {
	style := Styles[t.style]
	str := style[LineLEdge[typ]]
	for n := range t.columnsName {
		for i := 0; i < t.columnWidth(n); i++ {
			str += style[9]
		}

		if n == len(t.columnsName)-1 {
			str += style[LineREdge[typ]]
		} else {
			str += style[LineSplit[typ]]
		}
	}
	return str
}

func (t *Table) rowData(strs []string) string {
	style := Styles[t.style]
	str := style[10]

	for n, name := range strs {
		widthLen := t.columnWidth(n)
		runeLen := len([]rune(name))
		spaceRNum := (widthLen - runeLen) / 2
		spaceLNum := (widthLen - runeLen) - spaceRNum
		tmp1 := fmt.Sprintf("%-*v", runeLen+spaceRNum, name)
		tmp2 := fmt.Sprintf("%*v", spaceLNum+runeLen+spaceRNum, tmp1)
		str += tmp2
		str += style[10]
	}
	return str
}

func (t *Table) columnWidth(n int) int {
	v := t.columnsMaxLen[n] + width
	if v%2 != 0 {
		v++
	}
	return v
}
