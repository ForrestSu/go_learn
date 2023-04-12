package main

import (
	"fmt"

	"github.com/ForrestSu/go_learn/utils"
	"github.com/distatus/battery"
)

func main() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Failed to get battery info:", err)
		return
	}
	for _, b := range batteries {
		health := fmt.Sprintf("%0.2f%%", b.Full/b.Design*100)
		fmt.Printf("Battery: MaxCapcity %.f, health: %s\n", b.Full/b.Voltage, utils.TitlePt.Sprint(health))
		// fmt.Printf("%.f, %.f", b.Full/b.Voltage, b.Design/b.Voltage)
	}
}
