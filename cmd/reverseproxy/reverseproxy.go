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

//CONFIG is yaml config
const CONFIG string = "config.yml"

// loging server connection
var ConnTrace net.Conn

func main() {

	var err error

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
	cport := Proxy.Cport
	Servport := Proxy.Servers[0].Port

	fmt.Printf("Proxy Hostname:%v Port:%v", host, port)
	listener, _ := net.Listen("tcp", ":"+port)

	Cchan := make(chan string)
	for {
		go Start(listener, Cchan, Servport, cport)
		<-Cchan
	}

}

//Start the reverse proxy connection
func Start(listener net.Listener, Cchan chan string, Servport string, cport string) {
	conn, _ := listener.Accept()
	var serverConn net.Conn
	var err error
	// defer conn.Close()
	Cchan <- "Connection Esablished\n"

	buffer := make([]byte, 1024)
	conn.Read(buffer)

	for {

		serverConn, err = net.Dial("tcp", ":"+Servport)

		serverConn.Write(buffer)

		if err != nil {
		} else { // if error is nil we break out of loop
			break
		}

		if serverConn != nil { // If serverConn was found we break out of forever loop
			break
		}

	}
	go connTrace(cport)
	//New Variable InboundMessages a channel with string type
	InboundMessages := make(chan string)
	//New Variable OutboundMessages a channel with string type
	OutboundMessages := make(chan string)
	in := "INBOUND"
	out := "OUTBOUND"
	ins := "INBOUND SERVER"
	outs := "OUTBOUND SERVER"

	go Writer(conn, OutboundMessages, out)

	go Writer(serverConn, InboundMessages, ins)

	go Listener(serverConn, OutboundMessages, outs)

	go Listener(conn, InboundMessages, in)

	<-InboundMessages

}

func Writer(Conn1 net.Conn, messages chan string, in string) {

	for {
		//Set your channel message value as new message
		NewMessage := <-messages
		// If loging server exsists
		if ConnTrace != nil {
			//log new messeage
			x := NewMessage + " " + in
			ConnTrace.Write([]byte(x))
		}
		//Write channle set to connection passed in
		Conn1.Write([]byte(NewMessage))
	}
}

func Listener(Conn1 net.Conn, messages chan string, out string) {
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
		if ConnTrace != nil {
			// log message
			x := string(buf) + " " + out + "  "
			ConnTrace.Write([]byte(x))

		}
		//Send the buffer to message channle esablishing Read channel.
		messages <- string(buf)
	}

}
func connTrace(cport string) {
	var err error

	for {
		ConnTrace, err = net.Dial("tcp", ":"+cport)

		if err == nil {
			break
		}

		if ConnTrace != nil {
			break
		}

	}
}
