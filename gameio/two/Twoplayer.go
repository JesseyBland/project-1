// Package two houses the Two player option when -tp flag is used. X and O are user
//inputs.
package two

import (
	"fmt"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gameio/player"
	"github.com/JesseyBland/project-0/gamewin"
)

// Twoplayer this is the twopalyer routine
func Twoplayer() {
Reset:
	again := ""
	gamewin.CheckWin()
	for gamewin.CheckWin() == false {
		fmt.Print(gameboard.PrintBoard())
		player.Xmove()
		gamewin.CheckWin()

		if gamewin.CheckWin() == false {
			fmt.Print(gameboard.PrintBoard())
			twomove()
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
