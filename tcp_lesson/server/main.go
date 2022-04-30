package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. 创建tcp地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	tcpl, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, _ := tcpl.AcceptTCP()
		go handleConnection(conn)
	}

}

func handleConnection(conn *net.TCPConn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + string(buf[0:n]))
		str := "收到了" + string(buf[0:n])
		conn.Write([]byte(str))
	}
}
