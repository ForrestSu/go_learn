package utils

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

// StructToMap 将 struct 值编码为 map
func StructToMap(data interface{}) map[string]string {
	reflectValue := reflect.ValueOf(data)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	reflectType := reflectValue.Type()
	kvs := make(map[string]string)
	for i := 0; i < reflectValue.NumField(); i++ {
		tag := reflectType.Field(i).Tag.Get("json")
		kvs[tag] = formatValue(reflectValue.Field(i))
		// fmt.Printf("%s: %s\n", tag, kvs[tag])
	}
	return kvs
}

// formatValue 格式化
// 1) 基本类型 => 字符串
// 2) 复合类型/slice => json
func formatValue(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		return val.String()
	default:
		if val.Kind() == reflect.Slice && val.IsNil() {
			return "[]"
		}
		data, err := json.Marshal(val.Interface())
		if err != nil {
			log.Printf("json.Marshal() err:%v! value:%v\n", err, val)
		}
		return string(data)
	}
}
