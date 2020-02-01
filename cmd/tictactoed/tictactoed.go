package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

func init() {
	gameboard.LoadCells("[", "]")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./cmd/tictactoed/web/index.html")
	t.Execute(w, nil)
	gameboard.LoadCells("[", "]")
	gamewin.ResetWin()
	Xm = 0
	Om = 0

}
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./cmd/tictactoed/web/index.html")
	t.Execute(w, nil)

}
func xturn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "X turn!")
}
func oturn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "O turn!")
}
func hboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p2><br>%v%v%v <br></p2>", gameboard.Board[0], gameboard.Board[1], gameboard.Board[2])
	fmt.Fprintf(w, "<p2>%v%v%v <br></p2>", gameboard.Board[3], gameboard.Board[4], gameboard.Board[5])
	fmt.Fprintf(w, "<p2>%v%v%v <br></p2>", gameboard.Board[6], gameboard.Board[7], gameboard.Board[8])

}

func main() {
	fmt.Println("Server Status:Listening Host:localhost Port:9999")
	http.HandleFunc("/", index)
	http.HandleFunc("/ttt", homepage)
	http.HandleFunc("/ttt1", playervsAi)
	http.HandleFunc("/ttt2", playervsPlayer)
	http.HandleFunc("/ttt3", aivsAi)

	err := http.ListenAndServe(":9999", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
