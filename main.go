package main

import (
	_ "image/png"

	"github.com/CodersSquad-Classes/final-challenge-pacteam/pacman"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(896, 768)
	ebiten.SetWindowTitle("Pacman by Pacteam")

	game, err := pacman.NewGame()

	if err != nil {
		panic(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
