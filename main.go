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
var coin *ebiten.Image

const (
	/* Screen settings */
	screenWidth  = 650
	screenHeight = 770
)
const (
	tileSize = 16
	tileXNum = 25
)

func init() {
	var err error
	level, _, err = ebitenutil.NewImageFromFile("assets/level.png")
	if err != nil {
		log.Fatal(err)
	}
	var err2 error
	coin, _, err2 = ebitenutil.NewImageFromFile("assets/coin.png")
	if err2 != nil {
		log.Fatal(err2)
	}
}

// *** Core Ebiten functions *** //
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Score: ")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 50)
	screen.DrawImage(level, op)

	op.GeoM.Scale(0.02, 0.02)
	op.GeoM.Translate(0, 200)
	screen.DrawImage(coin, op)

}

func (g *Game) Update() error {
	fmt.Println("update")

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman by Pacteam")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
