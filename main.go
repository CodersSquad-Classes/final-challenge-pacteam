package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	enemies []Enemy
}

var level *ebiten.Image
var tile *ebiten.Image
var coin *ebiten.Image
var ghost *ebiten.Image
var numEnemies int
var coordinates [500][2]int

const (
	/* Screen settings */
	screenWidth  = 700 //mitad tablero 325
	screenHeight = 750 //mitad tablero
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

	tile, _, err = ebitenutil.NewImageFromFile("assets/tile.png")
	if err != nil {
		log.Fatal(err)
	}
	var err2 error
	coin, _, err2 = ebitenutil.NewImageFromFile("assets/coin.png")
	if err2 != nil {
		log.Fatal(err2)
	}
	ghost, _, err2 = ebitenutil.NewImageFromFile("assets/ghost.png")
	if err2 != nil {
		log.Fatal(err2)
	}
}

// *** Core Ebiten functions *** //
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Score: ")

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(0, 50)
	for i := 0; i < 20; i++ {
		screen.DrawImage(tile, op)
		op.GeoM.Translate(35, 0)
	}
	op.GeoM.Reset()
	op.GeoM.Translate(0, 50)
	for i := 0; i < 20; i++ {
		screen.DrawImage(tile, op)
		op.GeoM.Translate(0, 35)
	}
	op.GeoM.Reset()
	op.GeoM.Translate(665, 50)
	for i := 0; i < 20; i++ {
		screen.DrawImage(tile, op)
		op.GeoM.Translate(0, 35)
	}
	op.GeoM.Reset()
	op.GeoM.Translate(0, 715)
	for i := 0; i < 20; i++ {
		screen.DrawImage(tile, op)
		op.GeoM.Translate(35, 0)
	}

	op.GeoM.Reset()
	op.GeoM.Scale(0.02, 0.02)
	op.GeoM.Translate(100, 200)
	screen.DrawImage(coin, op)

	for _, e := range g.enemies {
		op.GeoM.Reset()
		op.ColorM.Reset()
		op.GeoM.Scale(0.05, 0.05)
		op.GeoM.Translate(float64(e.xPos), float64(e.yPos))
		op.ColorM.Translate(e.color[0], e.color[1], e.color[2], e.color[3])
		screen.DrawImage(ghost, op)
	}

}

func (g *Game) Update() error {
	fmt.Println("update")

	for _, enemy := range g.enemies {
		enemy.moveRandom()
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	//var colors [4][4]int
	colors := [5][4]float64{{0, 209, 255, 0}, {30, 0, 210, 0}, {250, 250, 250, 0}, {0, 0, 250, 0}, {245, 0, 131, 0}}

	numEnemies = 4

	en := make([]Enemy, numEnemies)
	for i := 0; i < numEnemies; i++ {
		en[i] = Enemy{
			xPos:  300 + (i * 20),
			yPos:  300 + (i * 20),
			color: colors[i],
		}
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pacman by Pacteam")

	g := &Game{enemies: en}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
