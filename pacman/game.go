package pacman

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	scene *scene
}

const (
	ScreenWidth  = 650
	ScreenHeight = 720
)

var (
	tiles_image *ebiten.Image
)

func NewGame() (*Game, error) {
	g := &Game{}

	var err error
	g.scene = createScene(nil)

	if err != nil {
		return nil, err
	}

	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	// if err := g.scene.Update(nil); err != nil {
	// 	return err
	// }

	return nil
}

var wall *ebiten.Image
var bg *ebiten.Image

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {

	wall, _, _ = ebitenutil.NewImageFromFile("assets/tile.png")
	bg, _, _ = ebitenutil.NewImageFromFile("assets/background.png")

	const xNum = ScreenWidth / tile_size

	for _, l := range g.scene.stage.tile_matrix {
		for i, t := range l {
			fmt.Print(string(t))
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tile_size), float64((i/xNum)*tile_size))

			// if current position is a 1, draw wall
			if string(t) == "1" {
				screen.DrawImage(wall, op)
			} else {
				screen.DrawImage(bg, op)
			}
		}
	}
}
