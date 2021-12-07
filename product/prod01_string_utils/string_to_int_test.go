package prod01_string_utils

import (
	"strconv"
	"testing"
)

func TestUint(t *testing.T) {
	str := "1152921504706122047"
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
}
