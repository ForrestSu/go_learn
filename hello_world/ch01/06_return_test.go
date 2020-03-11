package try_test

import (
	"os"
	"testing"
)

func TestReturn(t *testing.T) {
	t.Log("hello world!")
	os.Exit(-1)
}