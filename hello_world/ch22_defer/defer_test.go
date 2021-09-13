package ch22_defer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func handler() error {
	return fmt.Errorf("handler err is abandon")
}

func Filter() (err error) {
	var verifyErr error
	defer func() {
		err = verifyErr
		fmt.Printf("err = %+v, verifyErr = %+v \n", err, verifyErr)
	}()

	verifyErr = nil
	return handler()
}

func TestHello(t *testing.T) {
	err := Filter()
	assert.Nil(t, err) // need err
}
