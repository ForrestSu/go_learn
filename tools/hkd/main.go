package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/ForrestSu/go_learn/utils"
)

func main() {
	var link = "https://m.cmbchina.com/api/rate/fx-rate"
	var currency string
	flag.StringVar(&link, "url", link, "url")
	flag.StringVar(&currency, "currency", "港币", "币种名称: 美元")
	// 汇率金额换算
	var hkd, cny, rate float64
	flag.Float64Var(&hkd, "hkd", 0, "输入HKD")
	flag.Float64Var(&cny, "cny", 0, "输入CNY")
	flag.Float64Var(&rate, "r", 0, "指定汇率（默认: 实时招行中间价）")
	flag.Parse()
	// 如果没有指定汇率，则动态获取
	if rate == 0 {
		rate = getMidPrice(link, currency)
	}
	showOptions(hkd, cny, rate)
}

func showOptions(hkd float64, cny float64, rate float64) {
	rate = rate / 100.0
	if hkd > 0 {
		fmt.Printf("%v HKD => %.3f CNY\n", hkd, hkd*rate)
	}
	if cny > 0 {
		fmt.Printf("%v CNY => %.3f HKD\n", cny, cny/rate)
	}
}

func getMidPrice(url string, currency string) float64 {
	rsp, err := utils.HTTPGet(url)
	if err != nil {
		log.Fatal(err)
	}
	reply := &FxReply{}
	if err = json.Unmarshal(rsp, reply); err != nil {
		log.Fatalf("unmarshal err:%v", err)
	}
	reply.Print(currency)
	// 查询完成
	if currency == "" {
		return 0
	}
	// 查询指定币种汇率
	fxRate := reply.Find(currency)
	if fxRate == nil {
		log.Printf("币种未找到！currency=<%s>\n", currency)
		return 0
	}
	return fxRate.MidPrice()
}
