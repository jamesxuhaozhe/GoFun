package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write([]byte("hello")); err != nil {
		fmt.Println(err)
	}

	buffer := make([]byte, 1024)
	recvNum, err := conn.Read(buffer)

	msg := string(buffer[:recvNum])
	fmt.Println("recv from server: ", msg)

	req := []byte("hello")
	sendNum := 0
	num := 0
	// using a loop to make sure we have fully send the data
	for sendNum < len(req) {
		num, err = conn.Write(req[sendNum:])
		if err != nil {
			fmt.Println(err)
			break
		}
		sendNum += num
	}

	recvNum, err = conn.Read(buffer)
	msg = string(buffer[:recvNum])
	fmt.Println("recv from server: ", msg)
}
