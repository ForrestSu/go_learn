package main

import (
	_ "embed"

	"encoding/json"
	"fmt"
	"strconv"
)

//go:embed data_09-02.json
var data []byte

func main() {
	var power Power
	err := json.Unmarshal(data, &power)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total_power= %v\n", power.Data.TotalPower)
	var total float64
	for _, info := range power.Data.Infos {
		if info.Date >= "2024-08-26" {
			fmt.Printf("date=%s 中止计算\n", info.Date)
			break
		}
		fmt.Printf("计算%s = %s\n", info.Date, info.Power)
		dayPower, err := strconv.ParseFloat(info.Power, 64)
		if err != nil {
			panic(err)
		}
		total += dayPower
	}
	fmt.Printf("sum(power)= %f\n", total)
	fmt.Println("区间 [3298, 3434]=136度")
}

type Power struct {
	Data struct {
		Infos      []Info `json:"result"`
		TotalPower string `json:"totalPower"`
	} `json:"data"`
}

type Info struct {
	Date               string `json:"date"`
	MaxTemperature     string `json:"maxTemperature"`
	IsShareTime        string `json:"isShareTime"`
	MinTemperature     string `json:"minTemperature"`
	AverageTemperature string `json:"averageTemperature"`
	Power              string `json:"power"` // 度
	IsHot              string `json:"isHot"` // 是否是夏天
}
