package gameboard

import (
	"testing"
)

//Testing Each Cell Value
//Each Cell should be [i] and include a bool Fill thats false.
func TestCells(t *testing.T) {
	LoadCells("|", "|")
	testRgrid := "|"
	testLgrid := "|"
	testFill := false
	var testSlogic = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := range testSlogic {
		if Board[i].Slogic == testSlogic[i] && Board[i].RGrid == testRgrid && Board[i].LGrid == testLgrid && Board[i].Fill == testFill {
			t.Logf("Board %v Values Passed", i)
		} else {
			t.Errorf("Error Board %v Values Failed!", i)
		}

	}
}

// Testing the board printout sboard (which is the string representation of the board) with what it should be
// Initalized at when you Printboard without any user inputs.
func TestPrintboard(t *testing.T) {
	LoadCells("|", "|")

	testboard := "|1||2||3|\n|4||5||6|\n|7||8||9|\n"
	printboard := PrintBoard()
	if testboard == printboard {
		t.Logf("Pass")
	} else {
		t.Error("Fail")
	}
}
