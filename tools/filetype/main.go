package main

import (
	"flag"
	"fmt"

	"github.com/h2non/filetype"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "./1.bin", "二进制文件")
	flag.Parse()
	if kind, err := filetype.MatchFile(fileName); err == nil {
		fmt.Println("文件类型:", kind.Extension)
	} else {
		fmt.Printf("err=%v\n", err)
	}
}
