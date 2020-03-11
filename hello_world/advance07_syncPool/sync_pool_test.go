package advance07_syncPool

import (
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T){
	pool = &sync.Pool{
		New:nil,
	}
}