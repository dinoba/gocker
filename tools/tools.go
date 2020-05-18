package tools

import (
	"net"
	"os"
	"strings"
)

//GetIP returns first (not loopback) interface IP address.
func GetIP() string {
	//get IP
	myIP, _ := net.InterfaceAddrs()

	//fix IP 127.0.0.1
	myPublicIP := myIP[0].String()
	myPublicIPNoInterface := myPublicIP[:9]
	if myPublicIPNoInterface == "127.0.0.1" {
		myPublicIP = myIP[1].String()
		//remove /24
		myPublicIP = strings.Split(myPublicIP, "/")[0]
	}
	return myPublicIP
}

//GetHostname returns the host name reported by the kernel.
func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
