package try

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	origin := []int{1, 2, 3, 10, 11, 12}
	first := origin[0:3]
	second := []int{4, 5, 6}
	result := append(first, second...)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, result)
	fmt.Println("cap(first): ", cap(first)) // 输出：cap(first):  6
	fmt.Println("origin: ", origin)         // 输出：origin: [1 2 3 4 5 6]
	// 原数组origin的数据被修改了!
	// Append发现数组长度足够则不会分配一个新的slice，而是直接在原数组上修改。
	result = concatSlice(first, []int{1, 2, 3})
	fmt.Println("originV2: ", origin) // 输出: originV2: [1 2 3 4 5 6]

}

func concatSlice[T any](s1 []T, s2 []T) []T {
	n := len(s1)
	return append(s1[:n:n], s2...) // 限制 s1 的cap为n
}
