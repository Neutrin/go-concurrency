package main

import (
	"fmt"
	"net"
)

type Server interface {
	Start() error
}

type SimpleTcpServer struct {
	handler     func(conn net.Conn)
	workerLimit int64
	jobs        chan net.Conn
}

func NewSimpleTcpServer(workerLimit int64, hanlder func(conn net.Conn)) Server {
	return &SimpleTcpServer{
		handler:     hanlder,
		workerLimit: workerLimit,
		jobs:        make(chan net.Conn, 1000),
	}
}

func handleRequest(jobs <-chan net.Conn, handler func(conn net.Conn)) {
	for curJob := range jobs {
		handler(curJob)
	}
}

func (s *SimpleTcpServer) Start() error {
	listner, err := net.Listen("tcp", ":1729")
	if err != nil {
		return err
	}

	for count := 1; count <= int(s.workerLimit); count++ {
		go handleRequest(s.jobs, s.handler)
	}
	fmt.Println("Server started : started listening for request ")
	for {
		conn, err := listner.Accept()
		if err != nil {
			//If error comes close all the go routines created
			close(s.jobs)
			return err
		}
		fmt.Println(" request received")
		s.jobs <- conn

	}

}
