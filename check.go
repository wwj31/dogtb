package dogtb

import "reflect"

func checkAllInOne(arr ...interface{}) bool {
	var typ reflect.Type
	for _, v := range arr {
		if typ == nil {
			typ = reflect.TypeOf(v)
		}

		if typ != reflect.TypeOf(v) {
			return false
		}
	}
	return true
}
