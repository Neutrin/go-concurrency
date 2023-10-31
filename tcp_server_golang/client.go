package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func sendRequest() {
	fmt.Println(" request sent")
	curTime := time.Now()
	conn, err := net.Dial("tcp", "localhost:1729")

	defer conn.Close()
	conn.SetDeadline(time.Now().Local().Add(2 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte("hello"))
	if errors.Is(err, os.ErrDeadlineExceeded) {
		log.Println(" connection timed out")
		return
	}
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if errors.Is(err, os.ErrDeadlineExceeded) {
		log.Println(" connection timed out")
		return
	}

	fmt.Println(" Request completed resp ", string(buffer), " time taken = ", time.Since(curTime))
}

func callClient() {
	var wg sync.WaitGroup
	for count := 1; count <= 6; count++ {
		wg.Add(1)
		go sendRequest()
	}
	wg.Wait()
}
