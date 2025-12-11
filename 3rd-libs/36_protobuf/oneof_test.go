package oneof

import (
	"bytes"
	"fmt"
	"testing"

	pb "github.com/ForrestSu/go_learn/3rd-libs/36_protobuf/stub/echo"
	jsoniter "github.com/json-iterator/go"
	"github.com/kylelemons/godebug/pretty"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/encoding/protojson"
)

//go:generate protoc --go_out=. ./hello.proto

func TestOneOF(t *testing.T) {
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
}

func TestTmpMsg(t *testing.T) {
	// m := &pb.TmpMsg{
	// 	PreVid:    "sun",
	// 	WinPreVid: "quan",
	// 	MsgInfo:   "Hello",
	// }
	// data, _ := protojson.Marshal(m)
	// t.Log(string(data))

	// 解码
	jsonData := []byte(`{"previd":"V1","msg_info":"🙅🏻awesome"}`)
	recv := &pb.TmpMsg{}
	// case1： 如果使用trpc-http server 无需自己手动解码
	// case2：
	protojson.Unmarshal(jsonData, recv)
	t.Log(pretty.Sprint(recv))

	t.Log("============")
	jsonData2 := `{"pre_vid":"V2","msg_info":"🙅🏻awesome"}`
	recv2 := &pb.TmpMsg{}
	Unmarshaler.Unmarshal(bytes.NewBufferString(jsonData2), recv2)
	t.Log(pretty.Sprint(recv2))
}

var Unmarshaler = jsonpb.Unmarshaler{AllowUnknownFields: true}
