package main

import (
	"fmt"
	"strconv"

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

// Print 展示汇率
func (r *FxReply) Print(currency string) {
	if currency != "" {
		var got = r.Find(currency)
		if got != nil {
			fmt.Println(got.String())
		}
		return
	}
	fmt.Println("全部汇率:")
	for _, v := range r.Body.Data {
		fmt.Println(v.String())
	}
}

type FxRate struct {
	Name     string `json:"ccyNbr"` // 币种
	MidPx    string `json:"rtbBid"` // 中间价
	RthOffer string `json:"rthOfr"` // 现汇卖出(银行卖出)
	RthBid   string `json:"rthBid"` // 现汇买入
	RtcOffer string `json:"rtcOfr"` // 现钞卖出
	RtcBid   string `json:"rtcBid"` // 现钞买入
	Time     string `json:"ratTim"`
	Date     string `json:"ratDat"`
}

func (r *FxRate) String() string {
	return fmt.Sprintf(" %s MidRate: %s 现汇卖出: %s 现汇买入: %s (%s %s)",
		truncChinese(r.Name, 6), utils.TitlePt.Sprint(r.MidPx),
		utils.InfoPt.Sprint(r.RthOffer), r.RthBid, r.Date, r.Time)
}

func truncChinese(s string, width int) string {
	if len(s) >= width {
		return s[0:width]
	}
	return s
}

// MidPrice 汇率-中间价
func (r *FxRate) MidPrice() float64 {
	rate, _ := strconv.ParseFloat(r.MidPx, 10)
	return rate
}
