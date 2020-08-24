package main


/**
有时候我们需要导入一个package 来确保初始化 init()被调用过,
即使我们不需要使用 package 中任何函数或变量，使用下划线 _
 */

import (
	_ "github.com/ForrestSu/go_learn/hello_world/ch20_import_github/pkg/rectangle"
)

// rectangle.init() will be execute before main()
func main22() {
	println("enter pkg unused main...")
}
