package main

import (
	"fmt"
	"log"

	"github.com/kylelemons/godebug/pretty"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// 微信公众号-开发

var officialAccount *officialaccount.OfficialAccount

func InitWechat() {
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

	var redirectURL = "https://survey.video.qq.com/v1/h5/login?redirect_url=%2Fv1%2Fh5%2Fhello"

	url, _ := oauth.GetRedirectURL(redirectURL, "snsapi_userinfo", "state")
	log.Println(url)

	// oauth.GetUserAccessToken()

	// openID := "openID"
	// if userInfo, err := oauth.GetUserInfo(accessToken, openID, ""); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("userInfo = %v\n", userInfo)
	// }
}

// GetTemplateMsg 查询消息模板
func GetTemplateMsg() {
	var templates = officialAccount.GetTemplate()
	list, err := templates.List()
	log.Println(err)
	pretty.Print(list)
	// templates.Send()
}

// GetMaterial 查询所有图文消息
func GetMaterial() {
	var mater = officialAccount.GetMaterial()
	list, err := mater.BatchGetMaterial(material.PermanentMaterialTypeNews, 0, 20)
	log.Println(err)
	pretty.Print(list)

	var bc = officialAccount.GetBroadcast()
	result, err := bc.GetSpeed()
	log.Println(err)
	pretty.Print(result)
}

// PushTemplateMsg 发送模板消息
func PushTemplateMsg() {
	var template = officialAccount.GetTemplate()
	var templateID = "Li4GSGIxSkJC-wzk3YusSeh-SbtuOMwx2SjbeTpW21I"
	var pushMsg = &message.TemplateMessage{
		ToUser:     "o6A_q6pp-zdtd4UFmHI4Fl9aCaIY",
		TemplateID: templateID,
		// URL:        "",
		Data: map[string]*message.TemplateDataItem{
			"first":    {Value: "Hello"},      // 你好，会员卡序列码已送达！
			"keyword1": {Value: "math"},       // 会员卡序列码: GTestPAAAAA1234
			"keyword2": {Value: "100"},        // 会员卡类型: 会员月卡
			"keyword3": {Value: "20元"},        // 会员卡金额: 20元
			"keyword4": {Value: "2021年3月18日"}, // 时间: 2021年3月18日
			"remark":   {Value: "备注"},         // 备注: 请您前往-个人中心自助充值
		},
	}
	msgID, err := template.Send(pushMsg)
	// 1811836075779129347 <nil>
	log.Println(msgID, err)
}

func main() {
	InitWechat()
	// GetToken()
	// GetTemplateMsg()
	GetMaterial()
	PushTemplateMsg()
}
