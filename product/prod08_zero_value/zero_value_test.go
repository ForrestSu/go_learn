package prod08_zero_value

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 判断一个对象是否为值
func TestIsZeroValue(t *testing.T) {
	var integer int
	assert.True(t, reflect.ValueOf(integer).IsZero())

	var str string
	assert.True(t, reflect.ValueOf(str).IsZero())

	var list []int
	assert.True(t, reflect.ValueOf(list).IsZero())

	var ptr *int
	assert.True(t, reflect.ValueOf(ptr).IsZero())

	var flt float32
	assert.True(t, reflect.ValueOf(flt).IsZero())
}

//
// BenchmarkReflect
// BenchmarkReflect-4   	216344180	         5.511 ns/op
// BenchmarkPlain
// BenchmarkPlain-4     	1000000000	         0.3542 ns/op

func BenchmarkReflect(b *testing.B) {
	var integer int
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(integer).IsZero()
	}
}

func BenchmarkPlain(b *testing.B) {
	var integer int
	for i := 0; i < b.N; i++ {
		if integer == 0 {
			// ok
		}
	}
}
