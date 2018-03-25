package tagtry

import (
	"reflect"
)

func TagFunc(ptr interface{}) (fillds map[string]reflect.Value) {
	fillds = make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		filldInfo := v.Type().Field(i)
		tag := filldInfo.Tag
		name := tag.Get("json")
		if name == "" {
			continue
		}
		fillds[name] = v.Field(i)
	}
	return
}
