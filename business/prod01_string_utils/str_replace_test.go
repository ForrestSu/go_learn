package prod01_string_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAbc(t *testing.T) {
	str := "赵,钱,孙,李,赵"
	str1 := strings.Split(str, ",")
	fmt.Println(str1[0]) //赵
	fmt.Println(str1[1]) //钱
	fmt.Println(str1[2]) //孙
	fmt.Println(str1[3]) //李
	fmt.Println(str1[4]) //赵

	//字符串替换, -1表示全部替换, 0表示不替换, 1表示替换第一个, 2表示替换第二个...
	str2 := strings.Replace(str, "赵", "X", 1)
	fmt.Println(str2) //钱,钱,孙,李,钱

	abc := "https://baidu.com"
	//字符串截取, 一个汉字3个字节, 还有一个逗号
	str3 := abc[5:]
	fmt.Println(str3) // "://baidu.com"

}

func TestReplaceAnd(t *testing.T) {
	str := "ABCBC"
	actual := strings.ReplaceAll(str, "ABC", "A")
	assert.Equal(t, "ABC", actual)

	str2 := "hello world"
	actual2 := strings.ReplaceAll(str2, "hello", "one")
	assert.Equal(t, "one world", actual2)
}
