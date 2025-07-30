package main

import (
	"testing"
	"time"

	stub "github.com/ForrestSu/go_learn/3rd-libs/37_protobuf_any/stub/echo"
	"google.golang.org/protobuf/proto"
)

func TestCompact(t *testing.T) {
	var v2 = &stub.MsgV2{Name: "Bob", Extra: "新增字段"}
	v2Bytes, _ := proto.Marshal(v2)

	var v1 = &stub.MsgV1{}
	proto.Unmarshal(v2Bytes, v1)
	t.Logf("v1: %+v", v1)

	v1Bytes, _ := proto.Marshal(v1)

	// v1Msg -> v2Msg
	var v2Msg = &stub.MsgV2{}
	proto.Unmarshal(v1Bytes, v2Msg)
	t.Logf("v2: %+v", v2Msg)
}

func TestOK1(t *testing.T) {
	const expiredTimeSecond = 365 * 24 * (time.Hour / time.Second)
	t.Logf("expiredTimeSecond: %d", expiredTimeSecond)
}
