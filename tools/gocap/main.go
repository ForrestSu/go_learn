package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	capData()
}

func findDevices() {
	fmt.Println("----------Find all devices---------\n ")
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// Print device information
	for i, d := range devices {
		fmt.Printf("%d. %s\n", i+1, d.Name)
		// fmt.Printf("%v\n", d)
	}
}
