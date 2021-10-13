package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "hello world!")
	})

	if err := http.ListenAndServe("localhost:8082", nil); err != nil {
		log.Fatal(err)
	}
}
