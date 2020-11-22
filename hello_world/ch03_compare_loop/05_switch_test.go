package ch03_compare_loop

import (
	"fmt"
	"testing"
)

//Go Switch可以比较字符串，还支持比较多个case
//注意： go 里面是不需要 break
func TestMultiMatch(t *testing.T) {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

// 使用switch 代替if..else if，使代码更简洁
func TestEmptySwitch(t *testing.T) {
	for i := 1; i <= 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "Even")
		case i%2 == 1:
			t.Log(i, "odd")
		default:
			t.Log(i, "unknown")
		}
	}
}

func TestExpressionMatch(t *testing.T) {
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
}

/*
采用 fallthrough 可以贯穿匹配每个条件，
fallthrough 应该在一个case中最后声明，
如果它出现在中间的某个地方，编译器将抛出错误 fallthrough statement out of place
*/

func TestFallThrough(t *testing.T) {
	switch num := 2; { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	case num < 300:
		fmt.Printf("%d is lesser than 300\n", num)
	}
}
