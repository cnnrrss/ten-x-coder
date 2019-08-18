package main

import (
	"fmt"
	"net"
)

// Test with nc
// sudo nc -l 256 & go run sock.go

func main() {
	conn, err := net.Dial("tcp", "localhost:256")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("hello socket world"))
	if err != nil {
		fmt.Println(err)
	}
}
