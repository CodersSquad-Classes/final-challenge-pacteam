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
	scene         *scene
	mode          Mode
	enemies       []*Enemy
	player        *Pacman
	numEnemies    int
	score         int
	lives         int
	isSuperPilled bool
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
	height    = 0
	width     = 0
	sizeH     = 0
	sizeW     = 0
	gameFont  font.Face
	scoreFont font.Face
)

var wallSprite *ebiten.Image
var bgSprite *ebiten.Image
var pillSprite *ebiten.Image
var superPillSprite *ebiten.Image
var pacmanSprite *ebiten.Image
var ghostSprite *ebiten.Image

var enemyColors = [][4]float64{{-.60, .40, .0, 0}, {.5, .3, -.1, 0}, {.5, 0, 0, 0}, {0, -.1, .8, 0}, {.6, 0, 1, 0}, {-.7, .4, .8, 0}, {-.70, .4, .6, 0}}

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

	scoreFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(tileSize / 2),
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

	g.numEnemies = len(g.scene.enemyPositions)

	wallSprite, _, _ = ebitenutil.NewImageFromFile("assets/tile.png")
	bgSprite, _, _ = ebitenutil.NewImageFromFile("assets/background.png")
	pillSprite, _, _ = ebitenutil.NewImageFromFile("assets/dotSmall.png")
	superPillSprite, _, _ = ebitenutil.NewImageFromFile("assets/dotBig.png")
	pacmanSprite, _, _ = ebitenutil.NewImageFromFile("assets/pacman1.png")
	ghostSprite, _, _ = ebitenutil.NewImageFromFile("assets/ghostRed1.png")

	g.lives = 3
	g.score = 0

	height = len(g.scene.stage)
	width = len(g.scene.stage[0])

	sizeW = ((width*tileSize)/backgroundImageSize + 1) * backgroundImageSize
	sizeH = ((height*tileSize)/backgroundImageSize + 1) * backgroundImageSize

	g.player = &Pacman{
		sprite:  pacmanSprite,
		x:       g.scene.pacmanInitialX,
		y:       g.scene.pacmanInitialY,
		initX:   g.scene.pacmanInitialX,
		initY:   g.scene.pacmanInitialY,
		targetX: g.scene.pacmanInitialX,
		targetY: g.scene.pacmanInitialY,
		dir:     right,
		nextDir: right,
		game:    g,
	}

	return g
}

func initializeEnemies(g *Game) {
	g.enemies = make([]*Enemy, len(g.scene.enemyPositions))
	en := make([]*Enemy, g.numEnemies)
	for i := 0; i < g.numEnemies; i++ {
		en[i] = &Enemy{
			x:        g.scene.enemyPositions[i][0],
			y:        g.scene.enemyPositions[i][1],
			initialX: g.scene.enemyPositions[i][0],
			initialY: g.scene.enemyPositions[i][1],
			targetX:  g.scene.enemyPositions[i][0],
			targetY:  g.scene.enemyPositions[i][1],
			color:    enemyColors[i%len(enemyColors)],
			dir:      none,
			nextDir:  make(chan direction),
			stop:     make(chan struct{}),
			game:     g,
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

		if inpututil.IsKeyJustPressed(ebiten.KeyW) && g.numEnemies < len(g.scene.enemyPositions) {
			g.numEnemies += 1
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyS) && g.numEnemies > 1 {
			g.numEnemies -= 1
		}
	case ModeGame:
		for _, enemy := range g.enemies {

			enemy.move()

			if enemy.x/32 == g.player.x/32 && enemy.y/32 == g.player.y/32 {
				if g.isSuperPilled {
					enemy.reset()
					g.score += 200
				} else {
					g.playerDies()
					break
				}
			}

		}

		g.player.getInput()
		g.player.move()
	case ModeGameOver:

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = ModeMenu
			g.score = 0
			g.lives = 3
			g.scene.reset()
			g.player.reset()
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case ModeMenu:
		screen.Fill(color.Gray{0x7f})

		titleTexts := []string{"PACMAN by Pacteam"}
		texts := []string{"", "# of ENEMIES"}
		instructionsText := []string{"", "", "(w = +1, s = -1, space = START):"}
		enemiesText := []string{"", "", "", "", "", fmt.Sprint(g.numEnemies)}

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
	case ModeGame:
		// drawing background image
		for i := 0; i < sizeH/tileSize; i++ {
			y := float64(i * tileSize)

			for j := 0; j < sizeW/tileSize; j++ {
				options := &ebiten.DrawImageOptions{}

				x := float64(j * tileSize)

				options.GeoM.Translate(x, y)
				screen.DrawImage(bgSprite, options)
			}
		}

		// drawing the actual walls
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				options := &ebiten.DrawImageOptions{}

				x := float64(j * tileSize)
				y := float64(i * tileSize)

				options.GeoM.Translate(x, y)

				if g.scene.stage[i][j] == wall {
					screen.DrawImage(wallSprite, options)
				}

				if g.scene.stage[i][j] == pill {
					screen.DrawImage(pillSprite, options)
				}

				if g.scene.stage[i][j] == superPill {
					screen.DrawImage(superPillSprite, options)
				}
			}
		}

		// drawing the enemies
		for _, e := range g.enemies {
			e.Draw(screen, g.isSuperPilled)
		}

		g.player.draw(screen)

		//draw score and lives
		text.Draw(screen, fmt.Sprintf("Score: %v", g.score), scoreFont, 8, 24, color.White)
		text.Draw(screen, fmt.Sprintf("lives: %v", g.lives), scoreFont, 8, ScreenHeight-8, color.White)
	case ModeGameOver:
		// we're in the game over screen
		screen.Fill(color.Black)

		titleTexts := []string{"GAME OVER"}

		for i, l := range titleTexts {
			x := (ScreenWidth - len(l)*tileSize) / 24
			text.Draw(screen, l, gameFont, x, (ScreenHeight-tileSize)/2+tileSize*i, color.White)
		}

	}

}

func (g *Game) checkPill(i, j int) {
	switch g.scene.stage[i][j] {
	case pill:
		g.scene.stage[i][j] = empty
		g.score += 10
	case superPill:
		g.scene.stage[i][j] = empty
		g.score += 50
		go func() {
			g.isSuperPilled = true
			time.Sleep(20 * time.Second)
			g.isSuperPilled = false
		}()
	default:
		return
	}

	g.scene.remainingPills--
	if g.scene.remainingPills == 0 {
		g.player.reset()
		for _, enemy := range g.enemies {
			enemy.reset()
		}
		g.scene.reset()

	}
}

func (g *Game) playerDies() {
	g.lives--

	if g.lives == 0 {
		for _, enemy := range g.enemies {
			enemy.stopMovementAlgorithm()
		}
		g.mode = ModeGameOver

	} else {
		for _, enemy := range g.enemies {
			enemy.reset()
		}
	}
	g.player.reset()
	g.isSuperPilled = false

}
