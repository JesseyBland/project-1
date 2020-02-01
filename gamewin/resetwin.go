package gamewin

import "github.com/JesseyBland/project-0/gameboard"

// ResetWin resets the win states.
func ResetWin() {
	for i := range Win {
		Moves = 0
		WonChat = ""
		Win[i] = false
		gameboard.ResetBoard()

	}
}
