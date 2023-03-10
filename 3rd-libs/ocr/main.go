package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	// image to captcha
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./captcha.jpg")
	text, _ := client.Text()
	fmt.Println(text)
}
