package main

import (
	"log"

	"github.com/kylelemons/godebug/pretty"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var miniProgram *miniprogram.MiniProgram

// InitWechat 小程序初始化
func InitWechat() {
	memory := cache.NewMemory()
	wc := wechat.NewWechat()
	cfg := &miniConfig.Config{
		AppID:     "wx02de39eb8e3be430",
		AppSecret: appSecret,
		Cache:     memory,
	}
	miniProgram = wc.GetMiniProgram(cfg)
}

// Login 登陆
func Login() {
	var jsCode = "123"
	result, err := miniProgram.GetAuth().Code2Session(jsCode)
	if err != nil {
		log.Println(err)
		return
	}
	pretty.Print(result)
}

func main() {
	InitWechat()
	Login()
}
