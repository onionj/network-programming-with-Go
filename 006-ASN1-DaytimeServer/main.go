/* ASN1 DaytimeServer */

package main

import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Printf("Start a new listener on %s\n", tcpAddr.String())

	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("New Connection Accepted")

		daytime := time.Now()
		fmt.Println(daytime)

		mdata, _ := asn1.Marshal(daytime)

		conn.Write(mdata) // don't care about return value
		conn.Close()      // we're finished with this client
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
