package main

import (
	"fmt"
	"net"
	"time"
)

var ConnSignal chan string = make(chan string)
var LogConn net.Conn

func main() {

	port := ":3333"
	ln, _ := net.Listen("tcp", port)
	fmt.Printf("ConnTrace Status: UP\n")

	for {

		go Session(ln)
		fmt.Println(<-ConnSignal)

	}

}

func Session(ln net.Listener) {
	conn, _ := ln.Accept() //waits till it gets connection
	ConnSignal <- "ConnTrace Status:Connection"
	defer conn.Close()

	for {
		now := time.Now().Format(time.RFC850)
		buf := make([]byte, 1024)
		conn.Read(buf)
		fmt.Printf("%v\n", now)
		fmt.Print("  ↪Traffic Flow: " + string(buf) + "\n")
	}
}
