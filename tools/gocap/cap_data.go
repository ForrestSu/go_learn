package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func capData() {
	handle, err := pcap.OpenLive("en0", 65536, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer handle.Close()

	// Create a new PacketDataSource
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	// Packets returns a channel of packets
	in := src.Packets()

	for {
		var packet gopacket.Packet
		select {
		// case <-stop:
		// return
		case packet = <-in:
			arpLayer := packet.Layer(layers.LayerTypeIPv4)
			if arpLayer == nil {
				continue
			}
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer == nil {
				continue
			}
			tcp := tcpLayer.(*layers.TCP)
			fmt.Println("seq=", tcp.Seq)
			// .......

		}
	}
}
