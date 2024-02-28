package main

import (
	"flag"
	"fmt"
)

// 场景: 比如报销金额为5000, 凑满5000金额的最小的发票选取方案

type Bag struct {
	Value      int   // 当前背包的总价值
	SelectItem *Item // 最后选取的一个物品
}

type Solution struct {
	dp []Bag
}

func newSolution(maxValue int) *Solution {
	s := &Solution{dp: make([]Bag, maxValue+1)}
	for i := 0; i < len(s.dp); i++ {
		s.dp[i].Value = -1
	}
	s.dp[0].Value = 0
	return s
}

// ZeroOnePack 01背包: 每个物品选或者不选
func (s *Solution) ZeroOnePack(cost int, value int, item *Item) {
	maxValue := len(s.dp) - 1
	for V := maxValue; V >= cost; V-- {
		if s.dp[V-cost].Value >= 0 {
			valueIfSelected := s.dp[V-cost].Value + value
			if valueIfSelected > s.dp[V].Value {
				s.dp[V].Value = valueIfSelected
				s.dp[V].SelectItem = item
			}
		}
	}
}

// 递归打印方案
func (s *Solution) recursivePrintPlan(k int, V int) {
	if V > 0 {
		s.recursivePrintPlan(k+1, V-s.dp[V].SelectItem.Value)
		fmt.Printf("items[%d]: %s\n", k, s.dp[V].SelectItem.String())
	}
}

func main() {
	var targetValue, maxValue int
	flag.IntVar(&targetValue, "target", 888, "目标价值")
	flag.IntVar(&maxValue, "max", 920, "最大价值")
	flag.Parse()

	fmt.Printf("tickets size =%d\n", len(tickets))
	s := newSolution(maxValue)
	for _, item := range tickets {
		s.ZeroOnePack(item.Value, item.Value, item)
	}

	var isFind bool
	for i := targetValue; i <= maxValue; i++ {
		if s.dp[i].Value > 0 {
			fmt.Printf("money: %d, Minimum solution = %d.\n", targetValue, s.dp[i].Value)
			fmt.Printf("the plan is:\n")
			s.recursivePrintPlan(0, s.dp[i].Value)
			isFind = true
			break
		}
	}
	if isFind {
		fmt.Println("OK!")
	} else {
		fmt.Println("can't find the plan!")
	}
}
