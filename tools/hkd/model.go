package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ForrestSu/go_learn/utils"
)

type FxReply struct {
	Code     string `json:"returnCode"`
	ErrorMsg string `json:"errorMsg"`
	Body     struct {
		Data []*FxRate `json:"data"`
		Time string    `json:"time"` // 时间戳
	} `json:"body"`
}

func (r *FxReply) Find(name string) *FxRate {
	for _, v := range r.Body.Data {
		if v.Name == name {
			return v
		}
	}
	return nil
}

type FxRate struct {
	Name     string `json:"ccyNbr"` // 币种
	MidPx    string `json:"rtbBid"` // 中间价
	RthOffer string `json:"rthOfr"` // 现汇卖出(银行卖出)
	RtcOffer string `json:"rtcOfr"` // 现钞卖出
	RthBid   string `json:"rthBid"` // 现汇买入
	RtcBid   string `json:"rtcBid"` // 现钞买入
	Time     string `json:"ratTim"`
	Date     string `json:"ratDat"`
}

// 字符串替换器
var replacer = strings.NewReplacer("港币", "HKD")

func (r *FxRate) String() string {
	return fmt.Sprintf(" %s MidRate: %s 现汇卖出: %s (%s %s)",
		replacer.Replace(r.Name), utils.TitlePt.Sprint(r.MidPx), utils.InfoPt.Sprint(r.RthOffer), r.Date, r.Time)
}

// MidPrice 汇率-中间价
func (r *FxRate) MidPrice() float64 {
	rate, _ := strconv.ParseFloat(r.MidPx, 10)
	return rate
}
