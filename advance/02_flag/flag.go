package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "Tom", "Input your name")
var age = flag.Int("age", 18, "Input your age")
var f = flag.Bool("isVIP", false, "Is VIP")
var postCode int

func init() {
	flag.IntVar(&postCode, "postcode", 1234, "Input your post code")
}

func main() {
	flag.Parse()
	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
	fmt.Println("VIP:", *f)
	fmt.Println("postCode:", postCode)
}
