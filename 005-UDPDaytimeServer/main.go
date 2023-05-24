package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"

	UDPAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", UDPAddr)
	checkError(err)

	for {
		hanldeClient(conn)
	}

}

func hanldeClient(conn *net.UDPConn) {
	var buf [512]byte

	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Printf("%s:%s\n", addr.String(), string(buf[0:n]))

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
