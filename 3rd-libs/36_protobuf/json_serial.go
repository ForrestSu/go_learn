package main

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// JSONAPI is json packing and unpacking object, users can change
// the internal parameter.
var JSONAPI = jsoniter.ConfigCompatibleWithStandardLibrary

// Marshaler is jsonpb marshal object, users can change its params.
var Marshaler = jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}

// Unmarshaler is jsonpb unmarshal object, users can chang its params.
var Unmarshaler = jsonpb.Unmarshaler{AllowUnknownFields: true}

// JSONPBSerialization provides jsonpb serialization mode. It is based on
// protobuf/jsonpb.
type JSONPBSerialization struct{}

// Unmarshal deserialize the in bytes into body.
func (s *JSONPBSerialization) Unmarshal(in []byte, body interface{}) error {
	input, ok := body.(proto.Message)
	if !ok {
		return JSONAPI.Unmarshal(in, body)
	}
	return Unmarshaler.Unmarshal(bytes.NewReader(in), input)
}

// Marshal returns the serialized bytes in jsonpb protocol.
func (s *JSONPBSerialization) Marshal(body interface{}) ([]byte, error) {
	input, ok := body.(proto.Message)
	if !ok {
		return JSONAPI.Marshal(body)
	}
	var buf []byte
	w := bytes.NewBuffer(buf)
	err := Marshaler.Marshal(w, input)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
