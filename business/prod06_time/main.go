package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func nowTime() {
	/**
	时间戳（秒）：1605800945;
	时间戳（纳秒）：1605800945667282000;
	时间戳（毫秒）：1605800945667;
	时间戳（纳秒转换为秒）：1605800945;
	*/
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
}

func MD5Bytes(s string) string {
	ret := md5.Sum([]byte(s))
	return hex.EncodeToString(ret[:])
}

func main() {

	fmt.Println(MD5Bytes("123456"))
	// MD5 ("123456") = e10adc3949ba59abbe56e057f20f883e

}
