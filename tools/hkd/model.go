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
	Name   string `json:"ccyNbr"` // 币种
	MidPx  string `json:"rtbBid"` // 中间价
	RthOfr string `json:"rthOfr"` // 现汇卖出
	RtcOfr string `json:"rtcOfr"` // currency
	RthBid string `json:"rthBid"` // 现汇买入
	RtcBid string `json:"rtcBid"` // currency
	Time   string `json:"ratTim"`
	Date   string `json:"ratDat"`
}

// 字符串替换器
var replacer = strings.NewReplacer("港币", "HKD")

func (r *FxRate) String() string {
	return fmt.Sprintf(" %s Rate: %s (%s %s)",
		replacer.Replace(r.Name), utils.TitlePt.Sprint(r.MidPx), r.Time, r.Date)
}

// MidPrice 汇率-中间价
func (r *FxRate) MidPrice() float64 {
	rate, _ := strconv.ParseFloat(r.MidPx, 10)
	return rate
}
