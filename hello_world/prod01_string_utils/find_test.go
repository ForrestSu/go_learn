package prod01_string_utils

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestIndexFunc(t *testing.T) {

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))  // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
}
