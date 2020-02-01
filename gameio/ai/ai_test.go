package ai

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/JesseyBland/project-0/gameboard"
	"github.com/JesseyBland/project-0/gamewin"
)

//TestAiOmove tests O move logic loops with a n=5 to produce the 5 result across the board.
//Tests Variables: Moves int, Slogic, Fill bool.
func TestAiOmove(t *testing.T) {
	gameboard.LoadCells("[", "]")
	checkRepeat := true
	var n int
	for checkRepeat == true {
		n = 5
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
					fmt.Printf("O moves to %v\nTotal moves: %v\n", n, gamewin.Moves)
				}
			}
		}
	}
	if gameboard.Board[5].Slogic == "O" && gameboard.Board[5].Fill == true && gamewin.Moves == 1 {
		t.Logf("PASSED: Ai O Move succeeded!")
	} else {
		t.Errorf("FAILED: AI O Move out of bounds!")
	}
}

//TestAiXmove tests X move logic loops with a n=1 to produce the 1 result across the board.
//Tests Variables: Moves int, Slogic, Fill bool.
func TestAiXmove(t *testing.T) {
	gameboard.LoadCells("[", "]")
	checkRepeat := true
	var n int
	gamewin.Moves = 0
	for checkRepeat == true {
		n = 1
		if gamewin.Moves > 8 {
			checkRepeat = false
			break
		} else {
			for i := 0; i < len(gameboard.Board); i++ {
				if n == (i) && gameboard.Board[i].Fill == false {
					gameboard.Board[i].Slogic = "X"
					gameboard.Board[i].Fill = true
					gamewin.Moves++
					checkRepeat = false
					fmt.Printf("X moves to %v\nTotal moves: %v\n", n, gamewin.Moves)
				}
			}
		}
	}
	if gameboard.Board[1].Slogic == "X" && gameboard.Board[1].Fill == true && gamewin.Moves == 1 {
		t.Logf("PASSED: Ai X Move succeeded!")
	} else {
		t.Errorf("FAILED: AI X Move out of bounds! Value should be 1 S:%v F:%v M:%v", gameboard.Board[1].Slogic, gameboard.Board[1].Fill, gamewin.Moves)
	}
}

//TestRand test random n int ai logic. Should be 0-8 int.
func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var n int
	n = rand.Intn(8)
	if 0 <= n && n <= 8 {
		t.Logf("PASSED: Rand succeeded!")
	} else {
		t.Errorf("FAILED: n=%v Rand out of bounds!", n)
	}
}

//BenchmarkAivsAi Benchmarking My tictactoe AivsAi.
func BenchmarkAivsAi(b *testing.B) {
	// run the AivsAi inside the Reset loop b.N times
	gameboard.LoadCells("[", "]")
	gamewin.ResetWin()
	for n := 0; n < b.N; n++ {
		for gamewin.CheckWin() == false {
			fmt.Print(gameboard.PrintBoard())
			AiXmove()
			time.Sleep(1 * time.Second)
			gamewin.CheckWin()
			if gamewin.CheckWin() == false {
				fmt.Print(gameboard.PrintBoard())
				AiOmove()
				time.Sleep(1 * time.Second)
				gamewin.CheckWin()
			}
		}
		fmt.Print(gameboard.PrintBoard())
		fmt.Println("**************")
		fmt.Println(gamewin.WonChat)
	}
}
