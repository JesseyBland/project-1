// Package ai contains the ai move inputs
package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

//AiOmove is the O move in the Ai vs Ai.
func AiOmove() {
	checkRepeat := true
	var n int
	for checkRepeat == true {
		rand.Seed(time.Now().UnixNano())

		n = rand.Intn(9)
		if gamewin.Moves > 8 {
			checkRepeat = false
			break
		} else {
			for i := 0; i < len(gameboard.Board); i++ {
				if n == (i) && gameboard.Board[i].Fill == false {
					gameboard.Board[i].Slogic = "O"
					gameboard.Board[i].Fill = true
					gamewin.Moves++
					checkRepeat = false
					fmt.Printf("O moves to %v\nTotal moves: %v\n", n+1, gamewin.Moves)
				}
			}
		}

	}

}
