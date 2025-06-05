package t_test

import (
	"testing"

	"github.com/mitchellh/mapstructure"
)

type PageContext struct {
	Name string
	Age  int
	Sex  string
}

func TestMapToStruct(t *testing.T) {
	pageContext := &PageContext{}
	src := map[string]string{
		"Name": "test",
		"Age":  "18",
		"Sex":  "man",
	}
	err := mapstructure.Decode(src, pageContext)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pageContext)
}
