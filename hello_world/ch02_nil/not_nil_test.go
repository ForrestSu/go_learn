package niltest

import (
	"fmt"
	"os"
	"testing"
)

func Foo() error {
	var err *os.PathError = nil
	// …
	return err
}

func TestNil(t *testing.T) {
	err := Foo()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
	fmt.Printf("%#v\n", err)
}
