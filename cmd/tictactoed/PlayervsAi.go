package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gameio/ai"
	"github.com/JesseyBland/project-0/gamewin"
)

func playervsAi(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./cmd/tictactoed/web/ttt1.html")

	type Passdata struct {
		Cwin  bool
		Wonc  string
		Reset bool
		btn   string
	}
	c := Passdata{

		Cwin:  gamewin.CheckWin(),
		Wonc:  gamewin.WonChat,
		Reset: false,
		btn:   "Go!",
	}

	t.Execute(w, c)
	r.ParseForm()

	xread := strings.Join(r.Form["quantity"], "")
	x, _ := strconv.Atoi(xread)

	if x == 0 && gameboard.Board[x].Fill == false {
		hboard(w, r)

	} else {

		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				Xm++

				fmt.Printf("X moves to %v\nTotal moves: %v\n", x, gamewin.Moves)
				break
			}

		}

		gamewin.CheckWin()
		if gamewin.CheckWin() == false && Xm > Om {
			ai.AiOmove()
			Om++
			hboard(w, r)

			gamewin.CheckWin()

		} else {
			hboard(w, r)

		}
		if gamewin.CheckWin() == true && gamewin.WonChat != "O Wins!" {
			fmt.Fprintf(w, "<br><p1>%v</p1>", gamewin.WonChat)
			fmt.Println(gamewin.WonChat)

		} else if gamewin.CheckWin() == true && gamewin.WonChat == "O Wins!" {
			fmt.Fprintf(w, "<br><p1>%v</p1>", gamewin.WonChat)
			fmt.Println(gamewin.WonChat)

		}
	}
}
