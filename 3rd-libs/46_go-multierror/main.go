package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
)

func main() {
	err := multierror.Append(
		errors.New("error1"),
		errors.New("error2"),
		os.ErrNotExist,
	)
	fmt.Println(err.Error())
	fmt.Println(err.ErrorOrNil())
	fmt.Println(errors.Is(err, os.ErrNotExist))
}
