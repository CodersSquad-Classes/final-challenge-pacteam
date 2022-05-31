package pacman

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type direction int

const (
	none direction = iota
	up
	down
	right
	left
)

type Pacman struct {
	sprite           *ebiten.Image
	dir, nextDir     direction
	x, y             int
	targetX, targetY int
	game             *Game
}

func (p *Pacman) draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.sprite, op)
}

func (p *Pacman) move() {
	if p.x == p.targetX && p.y == p.targetY {
		p.nextTarget()
	}
	switch p.dir {
	case up:
		p.y--
	case down:
		p.y++
	case left:
		p.x--
	case right:
		p.x++
	}
}

func (p *Pacman) nextTarget() {

	if p.nextDir != none && !p.theresWall(p.nextDir) {
		p.dir = p.nextDir
	} else if p.theresWall(p.dir) {
		p.dir = none
	}

	switch p.dir {
	case up:
		p.targetY -= tileSize
	case down:
		p.targetY += tileSize
	case left:
		p.targetX -= tileSize
	case right:
		p.targetX += tileSize
	}
}

func (p *Pacman) theresWall(dir direction) bool {

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

	var i, j int
	i = (p.y + increaseY) / tileSize
	j = (p.x + increaseX) / tileSize

	if p.game.scene.stage.tile_matrix[i][j] == '#' {
		return true
	}
	return false
}
