package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

//Xm sets the X move number used in the logic for determining turn order
var Xm int

//Om sets the O move number used in the logic for determining turn order.
var Om int

//PlayervsPlayer is the two player option for my http tictactoe
func playervsPlayer(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./cmd/tictactoed/web/ttt2.html")
	var turn bool
	type Passdata struct {
		Cwin bool
		Wonc string
		Turn bool
		btn  string
	}
	c := Passdata{

		Cwin: gamewin.CheckWin(),
		Wonc: gamewin.WonChat,
		Turn: turn,
		btn:  "Go!",
	}

	t.Execute(w, c)
	r.ParseForm()

	xread := strings.Join(r.Form["Xquantity"], "")
	x, _ := strconv.Atoi(xread)

	oread := strings.Join(r.Form["Oquantity"], "")
	o, _ := strconv.Atoi(oread)
	if x == 0 && Xm == 0 {
		xturn(w, r)
		hboard(w, r)
	}

WC:
	if gamewin.CheckWin() == false {

		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false && Xm == Om {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				Xm++
				oturn(w, r)
				hboard(w, r)

				fmt.Printf("X moves to %v\nTotal moves: %v\n", x, gamewin.Moves)
				if gamewin.CheckWin() == true {
					goto WC
				}

			}

			if gamewin.CheckWin() == false {

				for i := 0; i < len(gameboard.Board); i++ {
					if o == (i+1) && gameboard.Board[i].Fill == false && Xm > Om {
						gameboard.Board[i].Slogic = "O"
						gameboard.Board[i].Fill = true
						gamewin.Moves++
						Om++

						xturn(w, r)
						hboard(w, r)

						fmt.Printf("O moves to %v\nTotal moves: %v\n", o, gamewin.Moves)
						if gamewin.CheckWin() == true {
							goto WC
						}

					}
				}

			}
		}

	} else if gamewin.CheckWin() == true {

		fmt.Fprintf(w, "<br><p1>%v</p1>", gamewin.WonChat)
		fmt.Println(gamewin.WonChat)

	}

}
