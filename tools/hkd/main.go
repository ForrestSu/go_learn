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
	var name = "港币"
	flag.StringVar(&link, "url", link, "url")
	var hkd, cny, rate float64
	flag.Float64Var(&hkd, "hkd", 0, "输入HKD")
	flag.Float64Var(&cny, "cny", 0, "输入CNY")
	flag.Float64Var(&rate, "r", 0, "指定汇率（默认: 实时招行中间价）")
	flag.Parse()

	if rate == 0 {
		rate = getMidPrice(link, name)
	}
	rate = rate / 100.0
	if hkd > 0 {
		fmt.Printf("%v HKD => %.3f CNY\n", hkd, hkd*rate)
	}
	if cny > 0 {
		fmt.Printf("%v CNY => %.3f HKD\n", cny, cny/rate)
	}
}

func getMidPrice(url string, name string) float64 {
	rsp, err := utils.HTTPGet(url)
	if err != nil {
		log.Fatal(err)
	}
	reply := &FxReply{}
	if err = json.Unmarshal(rsp, reply); err != nil {
		log.Fatalf("unmarshal err:%v", err)
	}
	fxRate := reply.Find(name)
	if fxRate == nil {
		log.Println("数据为空")
		return 0
	}
	fmt.Println(fxRate.String())
	return fxRate.MidPrice()
}
