package main


/**
有时候我们需要导入一个package 来确保初始化 init()被调用过,
即使我们不需要使用 package 中任何函数或变量，使用下划线 _
 */

import (
	_ "pkg/rectangle"
)

func Main1() {
	// rectangle.init() //unexported name
	println("enter pkg unused main...")
}
