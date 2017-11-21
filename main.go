package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 0 {
		fmt.Println("ip is empty")
		return
	}
	ipLong := os.Args[1]
	ipByte := make([]byte, 4)
	int10, err := strconv.Atoi(ipLong)
	if err != nil {
		fmt.Println("ip err")
		return
	}
	int32 := uint32(int10)
	binary.BigEndian.PutUint32(ipByte, int32)
	ip := net.IP(ipByte)
	fmt.Println(ip.String())
}
