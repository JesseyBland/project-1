package main

import "net"

import "fmt"

var ConnSignal chan string = make(chan string)
var connections []net.Conn

func main() {
	ln, _ := net.Listen("tcp", ":6060")

	for {
		go Session(ln)
		fmt.Println(<-ConnSignal)
	}

}

func Session(ln net.Listener) {
	conn, _ := ln.Accept() //waits till it gets connection
	ConnSignal <- "Connection Established"
	defer conn.Close()
	connections = append(connections, conn)

	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		for _, c := range connections {

			c.Write(buf)
		}
		//fmt.Println(string(buf))

	}
}
