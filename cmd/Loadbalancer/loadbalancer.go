package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"time"

	"gopkg.in/yaml.v2"
)

//tcpServers filled with our yaml config
type Proxy struct {
	Proxyhost string
	Proxyport string
	Servers   []Server
}

type Server struct {
	Hostname string
	Port     string
}

//CONFIG is yaml config
const CONFIG string = "config.yml"

// loging server connection
var logConn net.Conn

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
	//Proxy Hostname and Port set to port and host
	port := Proxy.Proxyport
	host := Proxy.Proxyhost

	//Cycling through servers to save the ports to Servports
	var Servports []string
	var Servhosts []string
	for i := range Proxy.Servers {
		Servports = append(Servports, Proxy.Servers[i].Port)
		Servhosts = append(Servhosts, Proxy.Servers[i].Hostname)
	}

	fmt.Printf("Proxy Hostname:%v Port:%v", host, port)
	listener, _ := net.Listen("tcp", ":"+port)

	Cchan := make(chan string)
	for {
		Servport := loadBalancer(Servports, Servhosts)
		go Start(listener, Cchan, Servport)
		<-Cchan
	}

}

//Start the reverse proxy connection
func Start(listener net.Listener, Cchan chan string, Servport string) {
	conn, _ := listener.Accept()
	var serverConn net.Conn
	var err error
	defer conn.Close()
	Cchan <- "Connection Esablished\n"

	buffer := make([]byte, 1024)
	conn.Read(buffer)

	for {

		serverConn, err = net.Dial("tcp", ":"+Servport)
		//defer serverConn.Close()
		serverConn.Write(buffer)

		if err != nil {
			fmt.Println(err)

		} else { // if error is nil we break out of loop
			break
		}

		if serverConn != nil { // If serverConn was found we break out of forever loop
			break
		}

	}

	//New Variable Shutdown a channel with string type
	Shutdown := make(chan string)
	//New Variable InboundMessages a channel with string type
	InboundMessages := make(chan string)
	//New Variable OutboundMessages a channel with string type
	OutboundMessages := make(chan string)

	go Writer(conn, OutboundMessages, Shutdown)

	go Writer(serverConn, InboundMessages, Shutdown)

	go Listener(serverConn, OutboundMessages, Shutdown)

	go Listener(conn, InboundMessages, Shutdown)
	// <- shutdown used to keep the forever loop from continually making go routines. A block
	<-Shutdown

}

func Writer(Conn1 net.Conn, messages chan string, shutdown chan string) {

	for {
		//Set your channel message value as new message
		NewMessage := <-messages
		// If loging server exsists
		if logConn != nil {
			//log new messeage
			logConn.Write([]byte(NewMessage))
		}
		//Write channle set to connection passed in
		Conn1.Write([]byte(NewMessage))
	}
}

func Listener(Conn1 net.Conn, messages chan string, shutdown chan string) {
	//Forever establish a reading connection
	for {
		//Create a Buffer with a meg size
		buf := make([]byte, 1024)
		//Read set to the buffer
		_, err := Conn1.Read(buf)
		//if error exsists
		if err != nil {
			//print error
			fmt.Println(err)
			//likelye error cause by server disconnect
			break
		}
		//if logging server exsists
		if logConn != nil {
			// log message
			logConn.Write(buf)
		}
		//Send the buffer to message channle esablishing Read channel.
		messages <- string(buf)
	}
	//Server shutdown.
	shutdown <- "Connection Closed"
}

func loadBalancer(Servports, Servhosts []string) string {
	var tn int
	for range Servhosts {
		tn++
	}
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(tn)
	now := time.Now()
	sp := Servports[x]
	sh := Servhosts[x]
	if logConn != nil {
		logConn.Write([]byte(" Loaded:" + sh + "Port:" + sp))
	}
	fmt.Printf("\n-Loaded: %v Port:%v    %v\n", sh, sp, now)
	return sp

}
