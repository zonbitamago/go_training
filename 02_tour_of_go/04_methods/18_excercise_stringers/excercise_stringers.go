package main

import (
	"fmt"
	"strings"
)

// IPAddr IPアドレス
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var str string
	for _, v := range ip {
		str += fmt.Sprintf("%v.", v)
	}

	str = strings.TrimRight(str, ".")

	return str
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
