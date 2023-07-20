package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var path = "."
	var localAddr = "localhost:8000"
	flag.StringVar(&path, "path", ".", "PATH")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(path)))
	fmt.Printf("visit http://%s ... path=%s\n", localAddr, path)
	if err := http.ListenAndServe(localAddr, nil); err != nil {
		log.Fatal(err)
	}
}
