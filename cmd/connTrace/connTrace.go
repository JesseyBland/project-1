package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"gopkg.in/yaml.v2"
)

//CONFIG is yaml config
const CONFIG string = "config.yml"

var ConnSignal chan string = make(chan string)
var LogConn net.Conn

//tcpServers filled with our yaml config
type Proxy struct {
	Proxyhost string
	Proxyport string
	Loadhost  string
	Loadport  string
	Chost     string
	Cport     string
	Servers   []Server
}

type Server struct {
	Hostname string
	Port     string
}

// loging server connection
var ConnTrace net.Conn

func main() {
	Proxy := Proxy{}

	file, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(file), &Proxy)
	if err != nil {
		panic(err)
	}
	//conntrace port pulled from config.yml
	cport := Proxy.Cport

	//Cycling through servers to save the ports to Servports

	ln, _ := net.Listen("tcp", ":"+cport)
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
