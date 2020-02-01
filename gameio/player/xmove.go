// Package player contains all the files regaurding player input vs ai output
package player

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

//X is x input
var X int

// Xmove is the X turn in the game.
func Xmove() {
	goodInput := false

	for goodInput == false {

		var user string
		if runtime.GOOS == "windows" {
			user = os.Getenv("USERNAME")
		} else {
			user = os.Getenv("USER")
		}

		fmt.Printf("%v's move pick 1-9\n", strings.Title(user))
		fmt.Scan(&X)

		for i := 0; i < len(gameboard.Board); i++ {
			if X == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.Moves++
				goodInput = true
				break
			}
		}

		if goodInput == false {
			fmt.Println("Error: The Number you entered has been chosen or you didn't enter a number.")
			fmt.Print(gameboard.PrintBoard())
		}
	}
}
