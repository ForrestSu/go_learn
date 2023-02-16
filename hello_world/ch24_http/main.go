package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ForrestSu/go_learn/utils"
)

type FxReply struct {
	Status int `json:"status"`
	Flag   int `json:"flag"`
	Data   []struct {
		Name  string `json:"ZCcyNbr"`
		MidPx string `json:"ZRtbBid"`
		Date  string `json:"ZRatDat"`
		Time  string `json:"ZRatTim"`
	} `json:"data"`
}

// 字符串替换器
var replacer = strings.NewReplacer("港币", "HKD")

func (r *FxReply) String() string {
	if len(r.Data) == 0 {
		return fmt.Sprintf("no-data! status=%d", r.Status)
	}
	d := r.Data[0]
	return fmt.Sprintf(" %s Rate: %s (%s %s)",
		replacer.Replace(d.Name), utils.TitlePt.Sprint(d.MidPx), d.Time, d.Date)
}

func (r *FxReply) HKDRate() float64 {
	rate, _ := strconv.ParseFloat(r.Data[0].MidPx, 10)
	return rate / 100
}

func main() {
	var url string
	flag.StringVar(&url, "url", "http://m.cmbchina.com/api/rate/getfxratedetail/?name=港币", "url")
	var hkd, cny float64
	flag.Float64Var(&hkd, "hkd", 0, "输入HKD")
	flag.Float64Var(&cny, "cny", 0, "输入CNY")
	flag.Parse()

	rsp, err := utils.HTTPGet(url)
	if err != nil {
		log.Fatal(err)
	}
	reply := &FxReply{}
	if err = json.Unmarshal(rsp, reply); err != nil {
		log.Fatalf("unmarshal err:%v", err)
	}
	fmt.Println(reply.String())
	if hkd > 0 {
		fmt.Printf("%v HKD => %.3f CNY\n", hkd, hkd*reply.HKDRate())
	}
	if cny > 0 {
		fmt.Printf("%v CNY => %.3f HKD\n", cny, cny/reply.HKDRate())
	}
}
