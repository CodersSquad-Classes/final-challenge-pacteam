package pacman

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	scene  *scene
	player *Pacman
}

const (
	ScreenWidth  = 896
	ScreenHeight = 768
)

var (
	height = 0
	width  = 0
	sizeH  = 0
	sizeW  = 0
)

var wall *ebiten.Image
var bg *ebiten.Image
var dotSmall *ebiten.Image
var dotBig *ebiten.Image
var pacman *ebiten.Image
var ghost *ebiten.Image

func NewGame() *Game {
	g := &Game{}

	g.scene = createScene(nil)

	wall, _, _ = ebitenutil.NewImageFromFile("assets/tile.png")
	bg, _, _ = ebitenutil.NewImageFromFile("assets/background.png")
	dotSmall, _, _ = ebitenutil.NewImageFromFile("assets/dotSmall.png")
	dotBig, _, _ = ebitenutil.NewImageFromFile("assets/dotBig.png")
	pacman, _, _ = ebitenutil.NewImageFromFile("assets/pacman1.png")
	ghost, _, _ = ebitenutil.NewImageFromFile("assets/ghostRed1.png")

	height = len(g.scene.stage.tile_matrix)
	width = len(g.scene.stage.tile_matrix[0])

	sizeW = ((width*tileSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH = ((height*tileSize)/backgroundImageSize + 1) * backgroundImageSize

	g.player = &Pacman{
		sprite:  pacman,
		x:       416,
		y:       448,
		targetX: 416,
		targetY: 448,
		dir:     right,
		nextDir: right,
		game:    g,
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.player.move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// drawing background image
	for i := 0; i < sizeH/tileSize; i++ {
		y := float64(i * tileSize)

		for j := 0; j < sizeW/tileSize; j++ {
			options := &ebiten.DrawImageOptions{}

			x := float64(j * tileSize)

			options.GeoM.Translate(x, y)
			screen.DrawImage(bg, options)
		}
	}

	// drawing the actual walls
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			options := &ebiten.DrawImageOptions{}

			x := float64(j * tileSize)
			y := float64(i * tileSize)

			options.GeoM.Translate(x, y)

			if g.scene.stage.tile_matrix[i][j] == '#' {
				screen.DrawImage(wall, options)
			}

			if g.scene.stage.tile_matrix[i][j] == '.' {
				screen.DrawImage(dotSmall, options)
			}

			if g.scene.stage.tile_matrix[i][j] == 'X' {
				screen.DrawImage(dotBig, options)
			}

			if g.scene.stage.tile_matrix[i][j] == 'G' {
				screen.DrawImage(ghost, options)
			}

			// if g.scene.stage.tile_matrix[i][j] == 'P' {
			// 	screen.DrawImage(pacman, options)
			// }
		}
	}

	g.player.draw(screen)
}
