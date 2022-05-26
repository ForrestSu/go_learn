package gob_vs_json_struct

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

/**
BenchmarkEncodeGob
BenchmarkEncodeGob-4               93618             11888 ns/op            3312 B/op         73 allocs/op
BenchmarkEncodeJson
BenchmarkEncodeJson-4             295291              4120 ns/op            1344 B/op          5 allocs/op
BenchmarkEncodeEasyJson
BenchmarkEncodeEasyJson-4        1227080               969 ns/op             960 B/op          4 allocs/op

BenchmarkDecodeGob
BenchmarkDecodeGob-4               30088             39045 ns/op           13029 B/op        356 allocs/op
BenchmarkDecodeJson
BenchmarkDecodeJson-4             201504              5973 ns/op             232 B/op          4 allocs/op
BenchmarkDecodeEasyJson
BenchmarkDecodeEasyJson-4         544663              2331 ns/op               0 B/op          0 allocs/op
PASS
*/

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

	jsonBytes, err := json.Marshal(origin)
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
		_, err := json.Marshal(origin)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}

// easyjson encoded
func BenchmarkEncodeEasyJson(b *testing.B) {
	origin := NewAidCache()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := origin.MarshalJSON()
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
	encoded, err := json.Marshal(origin)
	if err != nil {
		b.Errorf("err:%+v", err)
	}
	deObj := &AidCache{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(encoded, deObj)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}

// easyjson decoded
func BenchmarkDecodeEasyJson(b *testing.B) {
	origin := NewAidCache()
	encoded, err := origin.MarshalJSON()
	if err != nil {
		b.Errorf("err:%+v", err)
	}
	//////////
	deObj := &AidCache{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := deObj.UnmarshalJSON(encoded)
		if err != nil {
			b.Errorf("err:%+v", err)
		}
	}
}
