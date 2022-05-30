package pacman

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	scene   *scene
	enemies []Enemy
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

func NewGame(numEnemies int) *Game {
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

	colors := [8][4]float64{{0, 209, 255, 0}, {30, 0, 210, 0}, {0, 0, 0, 0}, {0, 0, 131, 0}, {0, 0, 131, 0}, {2, 2, 0, 0}, {0, 10, 0, 0}, {0, 5, 5, 0}}
	enemiesCoord := [8][2]int{{384, 320}, {416, 320}, {448, 320}, {480, 320}, {384, 352}, {416, 352}, {448, 352}, {480, 352}}
	en := make([]Enemy, numEnemies)
	for i := 0; i < numEnemies; i++ {
		en[i] = Enemy{
			xPos:  enemiesCoord[i][0],
			yPos:  enemiesCoord[i][1],
			color: colors[i],
		}
	}
	g.enemies = en

	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	for _, enemy := range g.enemies {
		enemy.moveRandom()
	}
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

			if string(g.scene.stage.tile_matrix[i][j]) == "#" {
				screen.DrawImage(wall, options)
			}

			if string(g.scene.stage.tile_matrix[i][j]) == "." {
				screen.DrawImage(dotSmall, options)
			}

			if string(g.scene.stage.tile_matrix[i][j]) == "X" {
				screen.DrawImage(dotBig, options)
			}
			/*
				if string(g.scene.stage.tile_matrix[i][j]) == "G" {
					screen.DrawImage(ghost, options)
					fmt.Print("x: ", x, " y: ", y, "\n")
				}*/
			if string(g.scene.stage.tile_matrix[i][j]) == "P" {
				screen.DrawImage(pacman, options)
			}
		}

	}

	//Draw enemies
	for _, e := range g.enemies {
		e.Draw(screen)
	}

}
