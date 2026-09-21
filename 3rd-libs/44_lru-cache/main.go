package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

func main() {
	cache := expirable.NewLRU[string, string](5, nil, time.Hour*10)
	v, ok := cache.Get("foo")
	if !ok {
		fmt.Println("key not found")
	} else {
		fmt.Println(v)
	}
	cache.Add("s1", "sun")
	cache.Add("s2", "moon")
	cache.Add("s3", "star")
	cache.Add("s4", "earth")
	cache.Add("s5", "mars")
	cache.Add("s6", "jupiter")
	cache.Add("s4", "quan")
	fmt.Println(cache.Len())
	for _, item := range cache.Keys() {
		fmt.Println("== " + item)
	}
}
