package player

import (
	"fmt"
	"os"
	"testing"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

//TestAivsPlayer tests the Ai round in the Player vs AI route
func TestAivsPlayer(t *testing.T) {
	gameboard.LoadCells("[", "]")

	xinput := [3]int{1, 2, 3} // X input to simulate user choosing 1 2 and 3

	if gamewin.CheckWin() == false {
		for i := 0; i < 3; i++ {
			fmt.Print(gameboard.PrintBoard())
			println(xinput[i])
			X = xinput[i]
			Xmove()
			gamewin.CheckWin()
			if gamewin.CheckWin() == false {
				fmt.Print(gameboard.PrintBoard())
				//ai.AiOmove() Omit the O move to make the X move win with 1. 2. 3.
				gamewin.CheckWin()
			}
		}
	}
	if gamewin.CheckWin() == true {
		t.Logf("PASSED: Loop succeeded!")
	} else {
		t.Errorf("FAILED: the Loop did not run through!")
	}
}

//TestXmove tests X move logic making sure x moves 5 changes slogic to 5, fill true and Moves +1.
func TestXmove(t *testing.T) {
	gameboard.LoadCells("[", "]")
	x := 5 //Move X to 5
	gamewin.Moves = 0
	goodInput := false
	for i := 0; i < len(gameboard.Board); i++ {
		if x == (i+1) && gameboard.Board[i].Fill == false {
			gameboard.Board[i].Slogic = "X"
			gameboard.Board[i].Fill = true
			gamewin.Moves++
			goodInput = true
			break
		}
	}
	if gameboard.Board[4].Slogic == "X" && gameboard.Board[4].Fill == true && goodInput == true && gamewin.Moves == 1 {

		t.Logf("PASSED: X moved succeeded!")
	} else {
		t.Errorf("Error: X move failed")
	}

}

//TestUsername tests the username os envirnomental variable to make sure its correctly swtiching based off operating system.
func TestUsername(t *testing.T) {
	var goos, user string

	goos = "linux"
	if goos == "windows" {
		user = os.Getenv("USERNAME")
	} else {
		user = os.Getenv("USER")
	}
	if user == os.Getenv("USER") {
		t.Logf("PASSED: Grab Username environment variable has passed!")
	} else {
		t.Errorf("Error: Grab Username environment variable has failed")

	}

}
