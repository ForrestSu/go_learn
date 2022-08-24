package try

import "testing"

type KeyValue struct {
	Key   string
	Value string
}

type Attributes []KeyValue

func TestHelle(t *testing.T) {
	var attrs []KeyValue
	var o Attributes = nil
	attrs = append(attrs, []KeyValue(o)...)
	t.Logf("attrs: %v  len(attrs)=%d", attrs, len(attrs))
}
