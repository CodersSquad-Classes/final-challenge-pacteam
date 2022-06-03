package pacman

import (
	"image/color"

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
	scene *scene
	mode  Mode
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
	height   = 0
	width    = 0
	sizeH    = 0
	sizeW    = 0
	gameFont font.Face
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
		Size:    float64(tileSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
}

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

	return g
}

func (g *Game) isKeyJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeMenu:
		if g.isKeyJustPressed() {
			g.mode = ModeGame
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	if g.mode == ModeMenu {
		// draw text in middle of screen
		titleTexts := []string{"PACMAN by Pacteam"}
		texts := []string{"", "", "", "", "", "", "", "PRESS SPACE KEY"}
		// screen.Draw(text, (ScreenWidth-w)/2, (ScreenHeight-w)/2, g.scene.font)

		for i, l := range titleTexts {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.White)
		}

		for i, l := range texts {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.White)
		}
	} else {
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

				if string(g.scene.stage.tile_matrix[i][j]) == "G" {
					screen.DrawImage(ghost, options)
				}

				if string(g.scene.stage.tile_matrix[i][j]) == "P" {
					screen.DrawImage(pacman, options)
				}
			}
		}
	}

}
