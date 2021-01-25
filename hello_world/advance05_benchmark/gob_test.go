package advance05_benchmark

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func init() {
	var data = &AidCache{}
	gob.Register(data)
}

// serialize
func serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// deserialize
func deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func TestGob(t *testing.T) {


}

func BenchmarkGob(b *testing.B) {

}
