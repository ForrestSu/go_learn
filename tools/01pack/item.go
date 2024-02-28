package main

import "fmt"

type Item struct {
	Value  int    // 价值
	Remark string // 备注
}

func (i *Item) String() string {
	return fmt.Sprintf("%s price %d", i.Remark, i.Value)
}

var tickets = []*Item{
	{Value: 903, Remark: "宜昌->广州南"},
	{Value: 388, Remark: "长沙->深圳北(二等)"},
	{Value: 890, Remark: "G1192虎门->信阳"},
	{Value: 103, Remark: "常德->长沙"},
}
