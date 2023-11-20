package atomic

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	var done atomic.Bool
	assert.False(t, done.Load())
	setTrue(&done)
	assert.True(t, done.Load())
}

func setTrue(done *atomic.Bool) {
	done.Store(true)
}
