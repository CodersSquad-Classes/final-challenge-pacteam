package pacman

import (
	"sync"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	sync.Mutex
	dir              direction
	nextDir          chan direction
	targetX, targetY int
	game             *Game
	xPos             int
	yPos             int
	color            [4]float64
}

func (e *Enemy) Draw(screen *ebiten.Image, g *Game) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(e.xPos), float64(e.yPos))
	options.ColorM.Translate(e.color[0], e.color[1], e.color[2], e.color[3])
	screen.DrawImage(ghost, options)
}
func (e *Enemy) travel() {
	for {
		dir := direction(rand.Intn(4) + 1)
		for i := rand.Intn(10) + 1; i > 0 && !e.theresWall(dir); i-- {
			e.nextDir <- dir
		}
	}

}
func (e *Enemy) updateTarget() {

	e.dir = <-e.nextDir
	e.Lock()
	switch e.dir {
	case up:
		e.targetY -= tileSize
	case down:
		e.targetY += tileSize
	case left:
		e.targetX -= tileSize
	case right:
		e.targetX += tileSize
	}
	e.Unlock()
}

func (e *Enemy) move() {
	if e.xPos == e.targetX && e.yPos == e.targetY {
		e.updateTarget()
	}
	switch e.dir {
	case up:
		e.yPos--
	case down:
		e.yPos++
	case left:
		e.xPos--
	case right:
		e.xPos++
	}
}
func (e *Enemy) theresWall(dir direction) bool {

	var increaseX, increaseY int

	switch dir {
	case up:
		increaseY -= tileSize
	case down:
		increaseY += tileSize
	case left:
		increaseX -= tileSize
	case right:
		increaseX += tileSize
	}

	e.Lock()
	var i, j int
	i = (e.targetY + increaseY) / tileSize
	j = (e.targetX + increaseX) / tileSize
	e.Unlock()

	if e.game.scene.stage.tile_matrix[i][j] == '#' {
		return true
	}
	return false
}
