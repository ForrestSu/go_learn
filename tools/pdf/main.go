package main

import (
	"flag"
	"fmt"
)

func main() {
	var usage = "mode模式：\n" +
		"remove_pwd: 移除密码 -in input.pdf -out output.pdf -pwd 123456\n" +
		"search_replace:  \n"
	var mode string
	flag.StringVar(&mode, "mode", "search_replace", usage)
	var input, output string
	flag.StringVar(&input, "in", "input.pdf", "输入文件")
	flag.StringVar(&output, "out", "output.pdf", "输出文件")
	// 移除密码
	var password string
	flag.StringVar(&password, "pwd", "", "密码")
	// 搜索替换
	var searchText, replaceText string
	flag.StringVar(&searchText, "search", "", "搜索文本")
	flag.StringVar(&replaceText, "replace", "", "替换文本")

	flag.Parse()

	switch mode {
	case "remove_pwd":
		if err := unlockPdf(input, output, password); err != nil {
			panic(err)
		}
		fmt.Printf("完成file: %s\n", output)
	case "search_replace":
		if err := RunSearchReplace(input, output, searchText, replaceText); err != nil {
			panic(err)
		}
		fmt.Printf("完成file: %s\n", output)
	default:
		fmt.Printf("不支持的操作: %s\n", output)
	}
}
