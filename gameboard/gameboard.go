// Package gameboard Creates a square Board filled with numbers each number surrounded by a string.
//Looking at cell 1 output: [|1 false|] : This is a struct with elements that define the left side | the right side | and
// the number 1, plus a bool variable saying if it is filled or empty true or false.
package gameboard

import (
	"fmt"
	"strconv"
)

//Board is a of structured cells.
var Board = make([]Cell, 0)

//row exsists to add a counter so that when it reaches scale it makes a new line for the next row.
var row = 0

//Cell the Cell structure of my board
type Cell struct {
	Slogic string
	LGrid  string
	RGrid  string
	Fill   bool
}

//This stringer will format for multiple scale levels. It should increase the cell size depending on the scale number of digits.
func (p Cell) String() string {
	switch {
	case Scale >= 4 && Scale <= 9:
		return fmt.Sprintf("%v%02v%v", p.LGrid, p.Slogic, p.RGrid)
	case Scale >= 10 && Scale <= 100:
		return fmt.Sprintf("%v%03v%v", p.LGrid, p.Slogic, p.RGrid)
	default:
		return fmt.Sprintf("%v%v%v", p.LGrid, p.Slogic, p.RGrid)
	}
}

//LoadCells fills the Cells structs with the board values.
func LoadCells(LGrid, RGrid string) {
	for i := 0; i < (Scale * Scale); i++ {
		Fill := false
		val := Cell{strconv.Itoa(i + 1), LGrid, RGrid, Fill}
		Board = append(Board, val)
	}
}

//ResetBoard Resets the board values to Cell number and Fill false.
func ResetBoard() {
	for i := 0; i < (Scale * Scale); i++ {
		Board[i].Fill = false
		Board[i].Slogic = strconv.Itoa(i + 1)
	}
}

//PrintBoard loops my board and prints the resulting cells. It handles the new line by useing a row counter based on the scale.
func PrintBoard() string {
	var sboard string
	for i := 1; i < (Scale*Scale)+1; i++ {

		row++
		if row == Scale {
			sboard += fmt.Sprintf("%v\n", Board[i-1])
			row = 0
		} else {
			sboard += fmt.Sprintf("%v", Board[i-1])
		}

	}
	return sboard
}
