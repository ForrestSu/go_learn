package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/kylelemons/godebug/pretty"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var officialAccount *officialaccount.OfficialAccount

func InitWechat() {
	/*
		https://api.weixin.qq.com/sns/userinfo?access_token=42_el5DMp-axQsZ2qhvoL-yh9tCuv0nFDD9sInND29psWnMdBn9_F-kTZJtAWwzIMUfio_rHdQAjcBBOA7vID3a2YSPsUrRu1gKjDQb2ntZnKtbo5_dKiVtN3-iMQbbqD8iyLTbayAE6al7m7KmTKZeAHAZQF&openid=wx39c38462bc2ca067&lang=zh_CN
	*/

	memory := cache.NewMemory()
	wc := wechat.NewWechat()
	cfg := &offConfig.Config{
		AppID:     "wx39c38462bc2ca067",
		AppSecret: appSecret,
		Token:     "xxx",
		Cache:     memory,
	}
	officialAccount = wc.GetOfficialAccount(cfg)
}

func GetToken() {
	oauth := officialAccount.GetOauth()
	accessToken, err := oauth.GetAccessToken()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("accessToken = %v\n", accessToken)
	}

	urlStr := url.QueryEscape("/v1/h5/hello")

	var redirectUrl = "https://survey.video.qq.com/v1/h5/login?redirect_url=%2Fv1%2Fh5%2Fhello"

	url, _ := oauth.GetRedirectURL(redirectUrl, "snsapi_userinfo", "state")
	log.Println(url)

	// oauth.GetUserAccessToken()

	//openID := "openID"
	//if userInfo, err := oauth.GetUserInfo(accessToken, openID, ""); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("userInfo = %v\n", userInfo)
	//}
}

func GetTemplateMsg() {
	var templates = officialAccount.GetTemplate()
	list, err := templates.List()
	log.Println(err)
	pretty.Print(list)
	// templates.Send()
}

func main() {
	InitWechat()
	GetToken()
	//GetTemplateMsg()
}
