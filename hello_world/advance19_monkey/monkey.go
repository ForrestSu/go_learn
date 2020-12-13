package main

import (
	"bou.ke/monkey"
	"fmt"
	"net/http"
	"reflect"
)

func main() {
	var d *http.Client // Has to be a pointer to because `Get()` has a pointer receiver
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Get", func(_ *http.Client, _ string) (*http.Response, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	// all http.Client instance, will be patched...
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
	fmt.Println()
}
