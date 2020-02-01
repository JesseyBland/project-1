package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gameio/player"
)

func main() {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", "localhost:6060")
		if err == nil { //
			break
		}
	}
	fmt.Println("Connection Established")
	fmt.Println("Press Enter to Start Tictactoe")
	go Listenter(conn)
	Writer(conn)
}

func Listenter(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		//fmt.Print(string(buf))

	}

}

func Writer(conn net.Conn) {
	for {
		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')

		conn.Write([]byte(text))
		gameboard.LoadCells("[", "]")
		player.Aiplayer()
		os.Exit(0)

	}

}
