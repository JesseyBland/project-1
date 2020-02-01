package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

//AiXmove is the X turn in the game played by ai.
func AiXmove() {
	checkRepeat := true
	var x int
	for checkRepeat == true {

		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(10)
		for i := 0; i < len(gameboard.Board); i++ {
			if x == (i+1) && gameboard.Board[i].Fill == false {
				gameboard.Board[i].Slogic = "X"
				gameboard.Board[i].Fill = true
				gamewin.CheckWin()
				gamewin.Moves++
				checkRepeat = false

			}

		}
	}
	gamewin.CheckWin()
	fmt.Printf("X moves to %v\nTotal moves: %v\n", x, gamewin.Moves)

}
