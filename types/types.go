package types

import (
	"fmt"
	"strconv"
	"strings"
)

//GetString 获取字符串
func GetString(v interface{}, def ...string) string {

	if v != nil {
		if vb, ok := v.([]byte); ok {
			return string(vb)
		}

		if r := fmt.Sprintf("%v", v); r != "" {
			return r
		}
	}
	if len(def) > 0 {
		return def[0]
	}

	return ""
}

//GetInt 获取int数据
func GetInt(v interface{}, def ...int) int {
	value := fmt.Sprintf("%v", v)
	if strings.Contains(strings.ToUpper(value), "E+") {
		var n float64
		_, err := fmt.Sscanf(value, "%e", &n)
		if err == nil {
			return int(n)
		}
		if len(def) > 0 {
			return def[0]
		}
	}
	if strings.Contains(value,".") {
		i := strings.Index(value,".")
		if value, err := strconv.Atoi(value[:i]); err == nil {
			return value
		}
		if len(def) > 0 {
			return def[0]
		}
	}

	if value, err := strconv.Atoi(value); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetInt64 获取int32数据，不是有效的数字则返回默然值或0
func GetInt32(v interface{}, def ...int32) int32 {

	if vb, ok := v.([]byte); ok {
		if i ,err :=strconv.Atoi(string(vb));err == nil {
			return int32(i)
		}
	}

	value := fmt.Sprintf("%v", v)
	if strings.Contains(strings.ToUpper(value), "E+") {
		var n float64
		_, err := fmt.Sscanf(value, "%e", &n)
		if err == nil {
			return int32(n)
		}
		if len(def) > 0 {
			return def[0]
		}
	}

	if strings.Contains(value,".") {
		i := strings.Index(value,".")
		if value, err := strconv.ParseInt(value[:i],10,32); err == nil {
			return int32(value)
		}
		if len(def) > 0 {
			return def[0]
		}
	}

	if value, err := strconv.ParseInt(value, 10, 32); err == nil {
		return int32(value)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetInt64 获取int64数据，不是有效的数字则返回默然值或0
func GetInt64(v interface{}, def ...int64) int64 {

	if vb, ok := v.([]byte); ok {
		if i ,err :=strconv.Atoi(string(vb));err == nil {
			return int64(i)
		}
	}

	value := fmt.Sprintf("%v", v)
	if strings.Contains(strings.ToUpper(value), "E+") {
		var n float64
		_, err := fmt.Sscanf(value, "%e", &n)
		if err == nil {
			return int64(n)
		}
		if len(def) > 0 {
			return def[0]
		}
	}

	if strings.Contains(value,".") {
		i := strings.Index(value,".")
		if value, err := strconv.ParseInt(value[:i],10,64); err == nil {
			return value
		}
		if len(def) > 0 {
			return def[0]
		}
	}

	if value, err := strconv.ParseInt(value, 10, 64); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetFloat32 获取float32数据，不是有效的数字则返回默然值或0
func GetFloat32(v interface{}, def ...float32) float32 {
	if value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 32); err == nil {
		return float32(value)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetFloat64 获取float64数据，不是有效的数字则返回默然值或0
func GetFloat64(v interface{}, def ...float64) float64 {
	if value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetUint32(v interface{}, def ...uint32) uint32 {
	value := fmt.Sprintf("%v", v)
	if strings.Contains(strings.ToUpper(value), "E+") {
		var n float64
		_, err := fmt.Sscanf(value, "%e", &n)
		if err == nil {
			return uint32(n)
		}
		if len(def) > 0 {
			return def[0]
		}
	}
	if value, err := strconv.ParseUint(value, 10, 32); err == nil {
		return uint32(value)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetUint64(v interface{}, def ...uint64) uint64 {
	value := fmt.Sprintf("%v", v)
	if strings.Contains(strings.ToUpper(value), "E+") {
		var n float64
		_, err := fmt.Sscanf(value, "%e", &n)
		if err == nil {
			return uint64(n)
		}
		if len(def) > 0 {
			return def[0]
		}
	}
	if value, err := strconv.ParseUint(value, 10, 64); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}
