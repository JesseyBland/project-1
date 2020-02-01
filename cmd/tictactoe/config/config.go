package config

import (
	"flag"

	"github.com/JesseyBland/project-0/gameboard"
)

//Av is the flag variable for AivsAi
var Av bool

//Tp is the flag variable for Two palyer
var Tp bool

//G1 is the first left and right pair selected grid options in gridconfig.txt file
var G1 bool

//G2 is the second left and right pair selected grid options in gridconfig.txt file
var G2 bool

//G3 is the third left and right pair selected grid options in gridconfig.txt file
var G3 bool

//G4 is the fourth left and right pair selected grid options in gridconfig.txt file
var G4 bool

//G5 is the fith left and right pair selected grid options in gridconfig.txt file
var G5 bool

//G6 is the sixth left and right pair selected grid options in gridconfig.txt file
var G6 bool

func init() {

	flag.BoolVar(&Tp, "tp", false, "Two Player Enabled")
	flag.BoolVar(&Av, "av", false, "Ai vs Ai Enable")
	flag.BoolVar(&G1, "g1", false, "Grid 1 Design")
	flag.BoolVar(&G2, "g2", false, "Grid 2 Design")
	flag.BoolVar(&G3, "g3", false, "Grid 3 Design")
	flag.BoolVar(&G4, "g4", false, "Grid 4 Design")
	flag.BoolVar(&G5, "g5", false, "Grid 5 Design")
	flag.BoolVar(&G6, "g6", false, "Grid 6 Design")
	flag.Parse()

	gameboard.LoadCells(assignGrid())

}

func assignGrid() (a, b string) {
	var lg string
	var rg string
	switch {
	case G1 == true:
		lg = gameboard.Grids[2]
		rg = gameboard.Grids[3]
		return lg, rg
	case G2 == true:
		lg = gameboard.Grids[4]
		rg = gameboard.Grids[5]
		return lg, rg
	case G3 == true:
		lg = gameboard.Grids[6]
		rg = gameboard.Grids[7]
		return lg, rg
	case G4 == true:
		lg = gameboard.Grids[8]
		rg = gameboard.Grids[9]
		return lg, rg
	case G5 == true:
		lg = gameboard.Grids[10]
		rg = gameboard.Grids[11]
		return lg, rg
	case G6 == true:
		lg = gameboard.Grids[12]
		rg = gameboard.Grids[13]
		return lg, rg
	default:
		lg = gameboard.Grids[0]
		rg = gameboard.Grids[1]
		return lg, rg
	}

}
