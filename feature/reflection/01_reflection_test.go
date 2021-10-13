package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

type order struct {
	orderId    int
	customName string
}

func reflectStruct(p interface{}) {
	// 获取类型
	var tp = reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	fmt.Println("Type ", tp)
	fmt.Println(v)

	k := tp.Kind()
	// Kind = struct Size = 24
	fmt.Println("Kind =", k, "Size =", tp.Size())

	// 如果为结构体类型
	if tp.Kind() == reflect.Struct {
		val := reflect.ValueOf(p)
		// 获取字段个数
		fmt.Println("field count:", val.NumField())
		// 遍历每一个字段
		for i := 0; i < val.NumField(); i++ {
			element := val.Field(i)
			fmt.Println("type:", element.Kind(), "value:", element)
			switch element.Kind() {
			case reflect.Int:
				//提取对应的值，为对应的类型
				id := element.Int()
				fmt.Printf("===> is_int %T %d \n", id, id)
			case reflect.String:
				str := element.String()
				fmt.Printf("===> is_string %T %s", str, str)
			default:
				fmt.Println("type is not support:", element.Kind())
			}
		}
	}
}

// Clear is better than clever.  Reflection is never clear.
func TestReflection(t *testing.T) {
	o := order{
		orderId:    10,
		customName: "Alice",
	}
	reflectStruct(o)
}
