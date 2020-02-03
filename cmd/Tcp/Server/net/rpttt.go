package main

import (
	"fmt"
	"io/ioutil"
	"net"

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
const CONFIG string = "./cmd/Tcp/Server/net/config.yml"

// loging server connection
var logConn net.Conn

func init() {

}

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
	port := Proxy.Proxyport
	host := Proxy.Proxyhost
	fmt.Printf("Proxy Hostname:%v Port:%v", host, port)
	listener, _ := net.Listen("tcp", ":"+port)

	Cchan := make(chan string)
	for {
		go Start(listener, Cchan, port)
		<-Cchan
	}

}

//Start the reverse proxy connection
func Start(listener net.Listener, Cchan chan string, port string) {
	conn, _ := listener.Accept()
	var serverConn net.Conn = nil
	var err error
	defer conn.Close()
	Cchan <- "Connection Esablished\n"

	buffer := make([]byte, 1024)
	conn.Read(buffer)
	Proxy := Proxy{}

	for {

		for i := range Proxy.Servers {
			ServerPrint := string(Proxy.Servers[i].Hostname + ":" + Proxy.Servers[i].Port)
			//if strings.Contains(string(buffer), ServerPrint) {
			serverConn, err = net.Dial("tcp", ServerPrint)
			defer serverConn.Close()
			serverConn.Write(buffer)
			//}
			if err != nil {
			} else { // if error is nil we break out of loop
				break
			}
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
