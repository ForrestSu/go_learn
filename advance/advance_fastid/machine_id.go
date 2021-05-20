package fastid

import (
	"net"
)

// GetMachineByIP 获取机器ID (IP最后2字节异或)
func GetMachineByIP(strIP string) byte {
	var ip = net.ParseIP(strIP)
	if len(ip) >= 2 {
		return ip[len(ip)-2] ^ ip[len(ip)-1]
	}
	return 0
}
