package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

const (
	/* Screen settings */
	screenWidth  = 540
	screenHeight = 540
)

func init() {
	fmt.Println("init")
}

// *** Core Ebiten functions *** //
func (g *Game) Draw(screen *ebiten.Image) {
	fmt.Println("draw")
}

func (g *Game) Update() error {
	fmt.Println("update")

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// *** *** //

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman by Pacteam")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
