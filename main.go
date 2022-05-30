package main

import (
	_ "image/png"

	"github.com/CodersSquad-Classes/final-challenge-pacteam/pacman"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(pacman.ScreenWidth, pacman.ScreenHeight)
	ebiten.SetWindowTitle("Pacman by Pacteam")

	game := pacman.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
