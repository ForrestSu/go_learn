package conc

import (
	"fmt"
	"testing"
	"time"

	"github.com/sourcegraph/conc"
)

func TestConcurrency(t *testing.T) {
	wg := conc.WaitGroup{}
	wg.Go(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("hello")
	})
	wg.Wait()
}
