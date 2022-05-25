package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

var level *ebiten.Image

const (
	/* Screen settings */
	screenWidth  = 650
	screenHeight = 720
)

func init() {
	var err error
	level, _, err = ebitenutil.NewImageFromFile("assets/level.png")
	if err != nil {
		log.Fatal(err)
	}
}

// *** Core Ebiten functions *** //
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(level, nil)
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
