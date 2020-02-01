// Package gamewin is a package that contains the games winning condiditons and functions
//that track win.
package gamewin

import (
	"github.com/JesseyBland/project-0/gameboard"
)

//Win is an array that houses the win states.
var Win [9]bool

//WonChat is a string array which outputs a Win or Lose message
var WonChat string

//Moves is a int that tracks number of X and O moves.
var Moves int

func wConditions() {

	switch {

	// X Win Conditions
	case gameboard.Board[0].Slogic == "X" && gameboard.Board[1].Slogic == "X" && gameboard.Board[2].Slogic == "X" && Win[0] == false:
		WonChat = "X Wins!"

		Win[0] = true

	case gameboard.Board[0].Slogic == "X" && gameboard.Board[3].Slogic == "X" && gameboard.Board[6].Slogic == "X" && Win[1] == false:
		WonChat = "X Wins!"

		Win[1] = true

	case gameboard.Board[0].Slogic == "X" && gameboard.Board[4].Slogic == "X" && gameboard.Board[8].Slogic == "X" && Win[2] == false:
		WonChat = "X Wins!"

		Win[2] = true

	case gameboard.Board[1].Slogic == "X" && gameboard.Board[4].Slogic == "X" && gameboard.Board[7].Slogic == "X" && Win[3] == false:
		WonChat = "X Wins!"

		Win[3] = true

	case gameboard.Board[2].Slogic == "X" && gameboard.Board[4].Slogic == "X" && gameboard.Board[6].Slogic == "X" && Win[4] == false:
		WonChat = "X Wins!"

		Win[4] = true

	case gameboard.Board[2].Slogic == "X" && gameboard.Board[5].Slogic == "X" && gameboard.Board[8].Slogic == "X" && Win[5] == false:
		WonChat = "X Wins!"

		Win[5] = true

	case gameboard.Board[3].Slogic == "X" && gameboard.Board[4].Slogic == "X" && gameboard.Board[5].Slogic == "X" && Win[6] == false:
		WonChat = "X Wins!"

		Win[6] = true

	case gameboard.Board[6].Slogic == "X" && gameboard.Board[7].Slogic == "X" && gameboard.Board[8].Slogic == "X" && Win[7] == false:
		WonChat = "X Wins!"

		Win[7] = true

	// O Win Conditions

	case gameboard.Board[0].Slogic == "O" && gameboard.Board[1].Slogic == "O" && gameboard.Board[2].Slogic == "O" && Win[0] == false:
		WonChat = "O Wins!"

		Win[0] = true

	case gameboard.Board[0].Slogic == "O" && gameboard.Board[3].Slogic == "O" && gameboard.Board[6].Slogic == "O" && Win[1] == false:
		WonChat = "O Wins!"

		Win[1] = true

	case gameboard.Board[0].Slogic == "O" && gameboard.Board[4].Slogic == "O" && gameboard.Board[8].Slogic == "O" && Win[2] == false:
		WonChat = "O Wins!"

		Win[2] = true

	case gameboard.Board[1].Slogic == "O" && gameboard.Board[4].Slogic == "O" && gameboard.Board[7].Slogic == "O" && Win[3] == false:
		WonChat = "O Wins!"

		Win[3] = true

	case gameboard.Board[2].Slogic == "O" && gameboard.Board[4].Slogic == "O" && gameboard.Board[6].Slogic == "O" && Win[4] == false:
		WonChat = "O Wins!"

		Win[4] = true

	case gameboard.Board[2].Slogic == "O" && gameboard.Board[5].Slogic == "O" && gameboard.Board[8].Slogic == "O" && Win[5] == false:
		WonChat = "O Wins!"

		Win[5] = true

	case gameboard.Board[3].Slogic == "O" && gameboard.Board[4].Slogic == "O" && gameboard.Board[5].Slogic == "O" && Win[6] == false:
		WonChat = "O Wins!"

		Win[6] = true

	case gameboard.Board[6].Slogic == "O" && gameboard.Board[7].Slogic == "O" && gameboard.Board[8].Slogic == "O" && Win[7] == false:
		WonChat = "O Wins!"
		Win[7] = true
	case Moves == 9 && WonChat != "X Wins!":
		WonChat = "No winners"
		Win[8] = true

	default:
		break
	}

}
