package reflection

import (
	"fmt"
	"reflect"
)

// func walk(x interface{}, fn func(string)) {
// 	val := reflect.ValueOf(x)
// 	field := val.Field(0)
// 	fn(field.String())
// }

func walk(x interface{}, fn func(input string)) {
	val := reflect.Indirect(reflect.ValueOf(x))

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	if val.Kind() == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		switch k := v.Kind(); k {
		case reflect.String:
			fn(v.String())
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				walk(val.Index(i).Interface(), fn)
			}
		case reflect.Struct:
			walk(v.Interface(), fn)
		default:
			fmt.Printf("skipping %v field", k)
		}

	}
}
