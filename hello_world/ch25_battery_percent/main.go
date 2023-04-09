package main

import (
	"fmt"

	"github.com/distatus/battery"
)

func main() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Failed to get battery info:", err)
		return
	}
	for _, b := range batteries {
		fmt.Printf("Battery: MaxCapcity %.f, %0.2f%%\n", b.Full/b.Voltage, b.Full/b.Design*100)
		// fmt.Printf("%.f, %.f", b.Full/b.Voltage, b.Design/b.Voltage)
	}
}
