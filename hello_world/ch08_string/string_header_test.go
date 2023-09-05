package try

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestStringPassByRef(t *testing.T) {
	str := "hello"
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("[ref] header=%p, data=%x, len=%d\n", header, header.Data, header.Len)
	// 底层指向的 string 数据是同一个 bytes
	assert.Equal(t, header.Data, stringParam(str))
}

func stringParam(s string) uintptr {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("[ref] header=%p, data=%x, len=%d\n", header, header.Data, header.Len)
	return header.Data
}

// empty string header.Data = nil
func Test_NilDataForEmptyString(t *testing.T) {
	var empty string
	header := (*reflect.StringHeader)(unsafe.Pointer(&empty))
	assert.True(t, header.Data == 0)
	assert.True(t, header.Len == 0)
}
