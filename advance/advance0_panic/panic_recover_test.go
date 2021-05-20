package advance0_panic_test

import (
	"fmt"
	"testing"
)

func test() {
	defer func() {
		// catch exception
		except := recover()
		fmt.Println(except)
	}()
	panic("test panic")
}

func TestRecover(t *testing.T) {
	test()
}
