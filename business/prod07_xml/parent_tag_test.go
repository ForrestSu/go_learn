package main

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * how to add parent Tags?
 *   use "a>b" struct label.
 */
type Service struct {
	XMLName xml.Name `xml:"service"`
	Port    int      `xml:"host>port"`
}

func TestParentTag(t *testing.T) {
	s := &Service{
		Port: 80,
	}
	expect := `<service><host><port>80</port></host></service>`
	actual, _ := xml.Marshal(s)
	assert.Equal(t, expect, string(actual))
}
