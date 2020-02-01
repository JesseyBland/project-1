package main

import (
	"github.com/JesseyBland/project-0/cmd/tictactoe/config"
	"github.com/JesseyBland/project-0/gameio/ai"
	"github.com/JesseyBland/project-0/gameio/player"
	"github.com/JesseyBland/project-0/gameio/two"
)

func main() {
	switch {
	case config.Tp == true:
		two.Twoplayer()
	case config.Av == true:
		ai.AivsAi()
	default:
		player.Aiplayer()
	}

}
