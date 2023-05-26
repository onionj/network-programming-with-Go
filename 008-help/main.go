/* just a test!! */

package main

import (
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
	checkError(err)
	fmt.Printf("Start a new listener on %s\n", tcpAddr.String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		fmt.Printf("(%s):(%s): Accepted\n", time.Now(), conn.RemoteAddr().String())

		go func(conn net.Conn) {
			defer conn.Close()

			response := `HTTP/1.1 200 OK
Content-Length: 8
Server: onionj
Content-Type: text/plain; charset=utf-8

Help Us!`
			conn.Write([]byte(response))
		}(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
