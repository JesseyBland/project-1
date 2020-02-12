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
	port := Proxy.Loadport
	host := Proxy.Loadhost
	cport := Proxy.Cport

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

		go Start(listener, Cchan, Servport, cport)
		<-Cchan
	}

}

//Start the reverse proxy connection
func Start(listener net.Listener, Cchan chan string, Servport string, cport string) {
	conn, _ := listener.Accept()
	var serverConn net.Conn
	var err error
	//defer conn.Close()
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
	go connTrace(cport)
	//New Variable InboundMessages a channel with string type
	InboundMessages := make(chan string)
	//New Variable OutboundMessages a channel with string type
	OutboundMessages := make(chan string)
	//Traffic direction variables for logs
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
		// If loging server exsistsgo connTrace()
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

			break
		}

		if ConnTrace != nil {
			// log message
			x := string(buf) + " " + out + "  "
			ConnTrace.Write([]byte(x))

		}

		messages <- string(buf)
	}

}

func loadBalancer(Servports, Servhosts []string) string {
	var tn int
	for range Servhosts {
		tn++
	}
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(tn)
	sp := Servports[x]
	sh := Servhosts[x]

	fmt.Printf("\n-Loaded: %v Port:%v    %v\n", sh, sp, now)
	return sp

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
