package main

import (
	"bou.ke/monkey"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func main() {
	MockFunc()
	MockInstance()
}

//  go run -gcflags=-l  .
func MockFunc() {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}

func MockInstance() {
	var d *http.Client // Has to be a pointer to because `Get()` has a pointer receiver
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Get", func(_ *http.Client, _ string) (*http.Response, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	// all http.Client instance, will be patched...
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
	fmt.Println()
}
