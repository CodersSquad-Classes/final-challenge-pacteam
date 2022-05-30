package pacman

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	scene *scene
}

const (
	ScreenWidth  = 896
	ScreenHeight = 768
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {

	return nil
}

var wall *ebiten.Image
var bg *ebiten.Image
var dotSmall *ebiten.Image
var dotBig *ebiten.Image
var pacman *ebiten.Image
var ghost *ebiten.Image

func (g *Game) Draw(screen *ebiten.Image) {

	wall, _, _ = ebitenutil.NewImageFromFile("assets/tile.png")
	bg, _, _ = ebitenutil.NewImageFromFile("assets/background.png")
	dotSmall, _, _ = ebitenutil.NewImageFromFile("assets/dotSmall.png")
	dotBig, _, _ = ebitenutil.NewImageFromFile("assets/dotBig.png")
	pacman, _, _ = ebitenutil.NewImageFromFile("assets/pacman1.png")
	ghost, _, _ = ebitenutil.NewImageFromFile("assets/ghostRed1.png")

	height := len(g.scene.stage.tile_matrix)
	width := len(g.scene.stage.tile_matrix[0])

	sizeW := ((width*tile_size)/background_image_size + 1) * background_image_size
	sizeH := ((height*tile_size)/background_image_size + 1) * background_image_size

	// drawing background image
	for i := 0; i < sizeH/tile_size; i++ {
		y := float64(i * tile_size)

		for j := 0; j < sizeW/tile_size; j++ {
			options := &ebiten.DrawImageOptions{}

			x := float64(j * tile_size)

			options.GeoM.Translate(x, y)
			screen.DrawImage(bg, options)
		}
	}

	// drawing the actual walls
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			options := &ebiten.DrawImageOptions{}

			x := float64(j * tile_size)
			y := float64(i * tile_size)

			options.GeoM.Translate(x, y)

			if string(g.scene.stage.tile_matrix[i][j]) == "#" {
				screen.DrawImage(wall, options)
			}

			if string(g.scene.stage.tile_matrix[i][j]) == "." {
				screen.DrawImage(dotSmall, options)
			}

			if string(g.scene.stage.tile_matrix[i][j]) == "X" {
				screen.DrawImage(dotBig, options)
			}

			if string(g.scene.stage.tile_matrix[i][j]) == "G" {
				screen.DrawImage(ghost, options)
			}

			if string(g.scene.stage.tile_matrix[i][j]) == "P" {
				screen.DrawImage(pacman, options)
			}
		}
	}
}
