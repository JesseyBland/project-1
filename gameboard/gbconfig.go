package gameboard

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

var Grids []string

var Scale int

//CONFIGFILE json file used to encode my text files into my project
const CONFIGFILE string = "./gameboard/boardconfig.json"

//Configuration struct that pulls my json files.
type Configuration struct {
	Grid        string `json:"Grid"`
	Scalenumber string `json:"Scalenumber"`
}

func init() {

	config := Configuration{}
	c, err := os.Open(CONFIGFILE)
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(c).Decode(&config)
	// Seting up the Left Gird and Right Grid
	g, err := os.Open(config.Grid)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()
	reader := bufio.NewScanner(g)
	for reader.Scan() {
		Grids = append(Grids, reader.Text())
	}
	// Setting up the Scale number.
	f, err := os.Open(config.Scalenumber)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader2 := bufio.NewScanner(f)
	var stringscale []string
	for reader2.Scan() {
		stringscale = append(stringscale, reader2.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
	Scale, err = strconv.Atoi(stringscale[0])
	if err != nil {
		log.Fatal(err)
	}

}
