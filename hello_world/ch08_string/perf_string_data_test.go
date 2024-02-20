package try

import (
	"unsafe"
)

// 直接通过强转string(bytes)或者[]byte(str)会带来数据的复制，性能不佳，所以在追求极致性能场景，
// 我们会采用『骇客』的方式，来实现这两种类型的转换,比如k8s采用下面的方式：
// https://github.com/kubernetes/apiserver/blob/706a6d89cf35950281e095bb1eeed5e3211d6272/pkg/authentication/token/cache/cached_token_authenticator.go#L263

// toBytes performs unholy acts to avoid allocations
func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// toString performs unholy acts to avoid allocations
func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 新版 如下函数签名：
//
// func String(ptr *byte, len IntegerType) string：根据数据指针和字符长度构造一个新的 string。
// func StringData(str string) *byte：返回指向该 string 的字节数组的数据指针。
// func SliceData(slice []ArbitraryType) *ArbitraryType：返回该 slice 的数据指针。
// func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType：根据数据指针和长度构造一个新的 slice。

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}
