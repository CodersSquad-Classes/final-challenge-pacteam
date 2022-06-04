package pacman

import (
	"sync"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	sync.Mutex
	dir                direction
	nextDir            chan direction
	stop               chan struct{}
	targetX, targetY   int
	x, y               int
	initialX, initialY int
	game               *Game
	color              [4]float64
}

func (e *Enemy) Draw(screen *ebiten.Image, g *Game) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(e.x), float64(e.y))
	options.ColorM.Translate(e.color[0], e.color[1], e.color[2], e.color[3])
	screen.DrawImage(ghostSprite, options)
}

func (e *Enemy) travel() {
	for {
		dir := direction(rand.Intn(4) + 1)
		for i := rand.Intn(5) + 1; i > 0 && !e.isWall(dir); i-- {
			select {

			case e.nextDir <- dir:
			case <-e.stop:
				return
			}
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
	if e.x == e.targetX && e.y == e.targetY {
		e.updateTarget()
	}
	switch e.dir {
	case up:
		e.y--
	case down:
		e.y++
	case left:
		e.x--
	case right:
		e.x++
	}
}

func (e *Enemy) isWall(dir direction) bool {

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

	if e.game.scene.stage[i][j] == wall {
		return true
	}

	return false
}

func (e *Enemy) reset() {
	e.stopMovementAlgorithm()
	e.x = e.initialX
	e.y = e.initialY
	e.targetX = e.initialX
	e.targetY = e.initialY
	e.dir = none
	go e.travel()
}

func (e *Enemy) stopMovementAlgorithm() {
	e.stop <- struct{}{}
}
