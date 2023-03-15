package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main1() {
	// Get a network listener on address "localhost:12345".
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", lis)

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
	})
	http.Serve(lis, nil)
}

type Adder interface {
	Add(context.Context, int, int) (int, error)
}
type adder struct {
	weaver.Implements[Adder]
}

func (adder) Add(_ context.Context, x, y int) (int, error) {
	return x + y, nil
}

//go:generate weaver generate .
func main() {
	ctx := context.Background()
	root := weaver.Init(ctx)
	adder, err := weaver.Get[Adder](root)
	if err != nil {
		log.Fatal(err)
	}
	sum, err := adder.Add(ctx, 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("1 + 2 = %d\n", sum)
}
