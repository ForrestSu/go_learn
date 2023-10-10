package go1_21

import (
	"sync/atomic"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	Name string
}

func TestAtomic(t *testing.T) {
	var gm atomic.Pointer[Config]
	v := &Config{Name: "hello"}
	gm.Store(v)
	assert.True(t, gm.Load() == v)
}

func TestZeroType(t *testing.T) {
	var zeroArray [0]int32
	assert.True(t, unsafe.Sizeof(zeroArray) == 0)
}
