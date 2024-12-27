package utils

import (
	"math/rand"
	"time"
)

type prizeItem struct {
	Kind int32
	Name string  // 奖品名称
	Prob float64 // 概率
}

var prizeList = []prizeItem{
	{Kind: 1, Prob: 0.5, Name: "一帆风顺卡"},
	{Kind: 3, Prob: 0.3, Name: "三阳开泰卡"},
	{Kind: 6, Prob: 0.15, Name: "六六大顺卡"},
	{Kind: 24, Prob: 0.05, Name: "十全十美卡"},
}

// 抽奖轮盘
func spinWheel(prizes []prizeItem, defVal int32) int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rv := rnd.Float64()
	sum := 0.0
	for _, item := range prizes {
		sum += item.Prob
		if rv < sum {
			return item.Kind
		}
	}
	panic("invalid prize list")
	return defVal
}
