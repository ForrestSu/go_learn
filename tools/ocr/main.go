package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "./captcha.jpg", "图片文件")
	flag.Parse()
	// check image file is exist
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("file <%s> not exist!\n", fileName)
		return
	}
	// image to captcha
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(fileName)
	text, err := client.Text()
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	fmt.Println(text)
}
