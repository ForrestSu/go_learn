package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserDate(t *testing.T) {
	str, _ := FormatTitle("10月(23-29) 62")
	assert.Equal(t, "2016-10-23 ~ 2016-10-29 (Week 62)", str)
}
