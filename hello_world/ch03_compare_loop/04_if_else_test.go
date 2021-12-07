package ch03_compare_loop

import (
	"fmt"
	"testing"
)

// 1 if 后面需要紧跟, 左花括号, 否则视为语法错
// 2 右花括号} 后面如果有else语句，} 和 else 必须在同一行，否则视为语法错
// 3 else 后面的花括号，可以写在下一行，也可以和 else 在同一行。
func TestIfElse(t *testing.T) {
	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}

// Go 支持两段的写法
func TestIF(t *testing.T) {
	if a := 1 + 1; a == 2 {
		t.Log("multi statement!")
	}
}

/**
原因是Go会自动插入分号。你可以在此处阅读有关分号插入规则的信息
https://golang.org/ref/spec#Semicolons。
如果这是该行的最后一个标记，Go会指定在 } 之后插入分号。因此，在if语句的 } 之后会自动插入分号。
所以我们的代码实际上是这样：

if num%2 == 0 {
      fmt.Println("the number is even")
};  //semicolon inserted by Go
else {
      fmt.Println("the number is odd")
}
*/

// 显式跳出-多重循环
func TestBreakOuter(t *testing.T) {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				println("break outer loop...")
				break outer
			}
		}
	}
	println("is break")
}
