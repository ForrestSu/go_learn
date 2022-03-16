package main

import (
	"fmt"
	"log"
)

var initVar int

// 再执行init()
func init() {
	initVar = 2
	log.Println("hello init()")
}

// GlobalVar 全局变量先初始化
var GlobalVar = newInit()

func newInit() *int {
	log.Println("hello NewInit()")
	var n = 1
	return &n
}

func main() {
	fmt.Printf("GlobalVar = %v, postCode:%d\n", *GlobalVar, initVar)
}
