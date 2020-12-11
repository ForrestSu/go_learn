package main

import (
	"encoding/xml"
	"fmt"
	"testing"
)

/**
 * 由于 go 不支持 self-closing tag！https://github.com/golang/go/issues/21399
 * 我们可以借助 innerxml 来做到这一点
 */

type Person struct {
	XMLName xml.Name `xml:"person"`
	SelfTag string   `xml:",innerxml"` // string 内容为 xml
	Comment string   `xml:",comment"`
}

func TestSelfClosingTag(t *testing.T) {
	person := &Person{
		Comment: "this is a comment!",
		SelfTag: "<inner/>",
	}
	output, _ := xml.Marshal(person)
	fmt.Println(string(output))
}
