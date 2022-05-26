package main

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/ForrestSu/go_learn/3rd-libs/36_protobuf/stub/echo"

	jsoniter "github.com/json-iterator/go"
)

//go:generate protoc --go_out=. ./hello.proto
func main() {
	req := &pb.HelloRequest{
		Msg:      "hello",
		OneOfPay: &pb.HelloRequest_Wx{Wx: &pb.WechatPay{Id: 100, Name: "wechat"}},
	}
	req.OneOfPay = &pb.HelloRequest_Noop{Noop: "Bob"}
	// val, err := json.MarshalIndent(req, "", "  ")
	// same as json.Marshal()
	// fmt.Println(JSONAPI.MarshalToString(req))

	// {"msg":"hello","OneOfPay":{"Wx":{"id":100,"name":"wechat"}}}
	fmt.Println(jsoniter.ConfigFastest.MarshalToString(req))

	val, err := protojson.Marshal(req)
	fmt.Printf("protojson.Marshal() => %s, err: %v\n", string(val), err)
	// jsonpb marshal
	var serial = JSONPBSerialization{}
	val, err = serial.Marshal(req)
	if err != nil {
		panic(err)
	}
	// {"msg":"hello","ali":{"id":"100","name":"ali"}}
	fmt.Printf("jsonpb.Marshal() => %s, err:%v\n", string(val), err)
}
