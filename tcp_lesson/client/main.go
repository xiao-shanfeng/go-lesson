package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	for {
		reader := bufio.NewReader(os.Stdin)
		bytes, _, _ := reader.ReadLine()
		_, _ = conn.Write(bytes)
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[0:n]))
	}
}
