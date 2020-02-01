package player

import (
	"fmt"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gameio/ai"
	"github.com/JesseyBland/project-0/gamewin"
)

//Aiplayer runs the game with ai controlling O.
func Aiplayer() {
Reset:
	var again string
	gamewin.CheckWin()
	for gamewin.CheckWin() == false {
		fmt.Print(gameboard.PrintBoard())
		Xmove()
		gamewin.CheckWin()
		if gamewin.CheckWin() == false {
			fmt.Print(gameboard.PrintBoard())
			ai.AiOmove()
			gamewin.CheckWin()
		}
	}
	fmt.Print(gameboard.PrintBoard())
	fmt.Println(gamewin.WonChat)
	fmt.Println("Play Again? y/n")

	fmt.Scan(&again)
	if again == "y" || again == "Y" || again == "Yes" || again == "yes" {

		gamewin.ResetWin()
		goto Reset

	}
}
