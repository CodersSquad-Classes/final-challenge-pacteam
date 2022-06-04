package pacman

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Mode int

type Game struct {
	scene   *scene
	mode    Mode
	enemies []*Enemy
	player  *Pacman
}

const (
	ScreenWidth  = 896
	ScreenHeight = 768
)

const (
	ModeMenu Mode = iota
	ModeGame
	ModeGameOver
)

var (
	height     = 0
	width      = 0
	sizeH      = 0
	sizeW      = 0
	numEnemies = 8
	gameFont   font.Face
)

var wall *ebiten.Image
var bg *ebiten.Image
var dotSmall *ebiten.Image
var dotBig *ebiten.Image
var pacman *ebiten.Image
var ghost *ebiten.Image

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)

	if err != nil {
		panic(err)
	}

	const dpi = 72
	gameFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(tileSize) - 5,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

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
	en := make([]*Enemy, numEnemies)
	for i := 0; i < numEnemies; i++ {
		en[i] = &Enemy{
			xPos:    enemiesCoord[i][0],
			yPos:    enemiesCoord[i][1],
			targetX: enemiesCoord[i][0],
			targetY: enemiesCoord[i][1],
			color:   colors[i],
			dir:     none,
			nextDir: make(chan direction),
			game:    g,
		}
		go en[i].travel()
	}
	g.enemies = en

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

func initializeEnemies(g *Game) {
	colors := [8][4]float64{{0, 209, 255, 0}, {30, 0, 210, 0}, {0, 0, 0, 0}, {0, 0, 131, 0}, {0, 0, 131, 0}, {2, 2, 0, 0}, {0, 10, 0, 0}, {0, 5, 5, 0}}
	enemiesCoord := [8][2]int{{384, 320}, {416, 320}, {448, 320}, {480, 320}, {384, 352}, {416, 352}, {448, 352}, {480, 352}}
	en := make([]*Enemy, numEnemies)
	for i := 0; i < numEnemies; i++ {
		en[i] = &Enemy{
			xPos:    enemiesCoord[i][0],
			yPos:    enemiesCoord[i][1],
			targetX: enemiesCoord[i][0],
			targetY: enemiesCoord[i][1],
			color:   colors[i],
			dir:     none,
			nextDir: make(chan direction),
			game:    g,
		}
		go en[i].travel()
	}

	g.enemies = en
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			initializeEnemies(g)
			g.mode = ModeGame
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyW) && numEnemies < 8 {
			numEnemies += 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyS) && numEnemies > 1 {
			numEnemies -= 1
		}
	case ModeGame:
		for _, enemy := range g.enemies {
			enemy.move()
		}

		g.player.getInput()
		g.player.move()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.mode == ModeMenu {
		screen.Fill(color.Gray{0x7f})

		titleTexts := []string{"PACMAN by Pacteam"}
		texts := []string{"", "# of ENEMIES"}
		instructionsText := []string{"", "", "(w = +1, s = -1, space = START):"}
		enemiesText := []string{"", "", "", "", "", fmt.Sprint(numEnemies)}

		for i, l := range titleTexts {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.Black)
		}

		for i, l := range texts {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.Black)
		}

		for i, l := range enemiesText {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.Black)
		}

		for i, l := range instructionsText {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.Black)
		}
	} else if g.mode == ModeGame {
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
			}
		}

		// drawing the enemies
		for _, e := range g.enemies {
			e.Draw(screen, g)
		}

		g.player.draw(screen)
	} else {
		// we're in the game over screen
		screen.Fill(color.Black)

		titleTexts := []string{"GAME OVER"}

		for i, l := range titleTexts {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.White)
		}

	}

}
