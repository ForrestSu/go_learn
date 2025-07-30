package main

import (
	"flag"
	"fmt"
	"os"

	stub "github.com/ForrestSu/go_learn/3rd-libs/37_protobuf_any/stub/echo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

//go:generate protoc --go_out=. ./hello.proto

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "3", "1 生成V2文件")
	flag.Parse()

	switch mode {
	case "1":
		var v2 = newMsgV2()
		v2Bytes, _ := proto.Marshal(v2)
		os.WriteFile("testdata/v2Bytes.txt", v2Bytes, 0644)
	case "2":
		// 中转服务v1（反序列化+再次序列化）
		v2Bytes, err := os.ReadFile("testdata/v2Bytes.txt")
		if err != nil {
			panic(err)
		}
		// 更新pb协议成旧版（没有新字段）
		var v1Msg = v2BytesToV1Msg(v2Bytes)
		v1Bytes, _ := proto.Marshal(v1Msg)
		if err = os.WriteFile("testdata/v1Bytes.txt", v1Bytes, 0644); err != nil {
			panic(err)
		}
	case "3":
		// v1Msg -> v2Msg
		v1Bytes, err := os.ReadFile("testdata/v1Bytes.txt")
		if err != nil {
			panic(err)
		}
		step3Show(v1Bytes)
	}
}

func step3Show(v1Bytes []byte) {
	var v2Msg = &stub.MsgV2{}
	proto.Unmarshal(v1Bytes, v2Msg)
	fmt.Printf("v2Msg: %+v\n", v2Msg)

	v2Ad := &stub.AdInfoV2{}
	if err := anypb.UnmarshalTo(v2Msg.GetAdInfo(), v2Ad, proto.UnmarshalOptions{}); err != nil {
		panic(err)
	}
	fmt.Printf("---> v2Ad: %+v\n", v2Ad)
}

// 第1步：构造 msgV2
func newMsgV2() *stub.MsgV2 {
	v2Ad := &stub.AdInfoV2{
		Id:    1,
		Title: "AdV2",
		// New:   "new_field",
	}
	anyData, _ := anypb.New(v2Ad)
	return &stub.MsgV2{
		Name:   "Alice",
		AdInfo: anyData,
		Extra:  "MsgNewField",
	}
}

// 第2步
func v2BytesToV1Msg(v2Bytes []byte) *stub.MsgV1 {
	var v1 = &stub.MsgV1{}
	proto.Unmarshal(v2Bytes, v1)
	fmt.Printf("v1Msg: %+v\n", v1)

	var v1Ad = &stub.AdInfoV2{}
	if err := anypb.UnmarshalTo(v1.GetAdInfo(), v1Ad, proto.UnmarshalOptions{}); err != nil {
		panic(err)
	}
	// 再次将 pbAny 反序列化回去
	if err := anypb.MarshalFrom(v1.AdInfo, v1Ad, proto.MarshalOptions{}); err != nil {
		panic(err)
	}
	return v1
}
