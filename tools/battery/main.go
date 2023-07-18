package main

import (
	"fmt"

	"github.com/ForrestSu/go_learn/utils"
)

func main() {
	// GetAll() copy from github.com/distatus/battery v0.11.0
	batteries, err := GetAll()
	if err != nil {
		fmt.Println("Failed to get battery info:", err)
		return
	}
	for i, b := range batteries {
		health := fmt.Sprintf("%.2f%%", 100*float64(b.MaxCapacity)/float64(b.DesignCapacity))
		fmt.Printf("<%d>: MaxCapcity %d (desgin: %d), "+
			"State: %s, "+
			"CycleCount: %s, "+
			"health: %s\n",
			i, b.MaxCapacity, b.DesignCapacity,
			utils.TitlePt.Sprint(b.State),
			utils.TitlePt.Sprint(b.CycleCount),
			utils.InfoPt.Sprint(health),
		)
	}
}
