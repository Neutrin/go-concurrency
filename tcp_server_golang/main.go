package main

import (
	"net"
	"time"
)

func processRequest(conn net.Conn) {
	data := make([]byte, 1024)
	defer conn.Close()
	conn.Read(data)
	time.Sleep(8 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK \r\n\r\nRequest Completed\r\n"))
}
func main() {
	server := NewSimpleTcpServer(5, processRequest)
	go server.Start()

	callClient()
}
