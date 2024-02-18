package main

import (
	"fmt"
	"os"

	"github.com/caiguanhao/readqr"
	"github.com/skip2/go-qrcode"
)

/*
qrcode.WriteFile()
1 content: 表示要生成二维码的内容，可以是任意字符串。
2 level: 表示二维码的容错级别，取值有Low、Medium、High、Highest。
3 size: 表示生成图片的width和height，像素单位。
4 filename: 表示生成的文件名(含路径)
*/
func main() {
	content := "https://forrestsu.github.io/"
	err := qrcode.WriteFile(content, qrcode.Medium, 256, "./qrcode.png")
	if err != nil {
		fmt.Println("fail to generate qrcode:", err)
	}
	decodeQrCode("./1.png")
}

func decodeQrCode(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result, err := readqr.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
