package advance05_benchmark

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
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
	// gob.Register(value)
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

func SerialJson(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}
func DeSerialJson(value []byte, v interface{}) error {
	return json.Unmarshal(value, v)
}

func TestGob(t *testing.T) {
	origin := NewAidCache()
	t.Logf("origin: %+v", origin)
	encoded, err := serialize(origin)
	// t.Logf("gob len %d, encoded: %s", len(encoded), string(encoded))
	assert.Nil(t, err)
	decoded, err := deserialize(encoded)

	val, ok := decoded.(*AidCache)
	assert.True(t, ok)
	assert.True(t, reflect.DeepEqual(origin, val))

	jsonBytes, err := SerialJson(origin)
	// t.Logf("json len %d, encoded: %s", len(jsonBytes), string(jsonBytes))
	_ = jsonBytes
}

// gob encoded
func BenchmarkEncodeGob(b *testing.B) {
	origin := NewAidCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := serialize(origin)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}

// json encoded
func BenchmarkEncodeJson(b *testing.B) {
	origin := NewAidCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := SerialJson(origin)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}

// decoded
func BenchmarkDecodeGob(b *testing.B) {
	origin := NewAidCache()
	encoded, err := serialize(origin)
	if err != nil {
		b.Errorf("err:%+v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := deserialize(encoded)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}

// json encoded
func BenchmarkDecodeJson(b *testing.B) {
	origin := NewAidCache()
	encoded, err := SerialJson(origin)
	if err != nil {
		b.Errorf("err:%+v", err)
	}
	deObj := &AidCache{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := DeSerialJson(encoded, deObj)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}
