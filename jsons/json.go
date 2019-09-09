package jsons

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Escape 把编码 \\u0026，\\u003c，\\u003e 替换为 &,<,>
func Escape(input string) string {
	r := strings.Replace(input, "\\u0026", "&", -1)
	r = strings.Replace(r, "\\u003c", "<", -1)
	r = strings.Replace(r, "\\u003e", ">", -1)
	r = strings.Replace(r, "\n", "", -1)
	return r
}

//Unmarshal 反序列化JSON
func Unmarshal(buf []byte) (c map[string]interface{}, err error) {
	c = make(map[string]interface{})
	err = json.Unmarshal(buf, &c)
	return
}

//Marshal 序列化JSON
func Marshal(v interface{}, tag ...string) (b []byte, err error) {

	tagName := "json"
	if len(tag) != 0 {
		tagName = tag[0]
	}

	switch tagName {
	case "json":
		return json.Marshal(v)
	default:
		t := reflect.TypeOf(v)
		val := reflect.ValueOf(v)
		mp := map[string]interface{}{}
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)
			v, ok := sf.Tag.Lookup(tagName)
			if ok {
				mp[v] = val.Field(i).Interface()
			}
		}
		return json.Marshal(mp)
	}
}
