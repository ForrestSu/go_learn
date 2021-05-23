package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockErr() (int, error) {
	return 100, fmt.Errorf("hello")
}

func caseShadow() (err error) {
	// 不会重新定义 err
	val, err := mockErr()
	log.Println(val, err)
	return
}

func caseNotBeShadow() (err error) {
	if true {
		// 重新定义 err
		val, err := mockErr()
		log.Println(val, err)
	}
	return
}

func TestShadow(t *testing.T) {
	assert.NotNil(t, caseShadow())
	assert.Nil(t, caseNotBeShadow())
}
