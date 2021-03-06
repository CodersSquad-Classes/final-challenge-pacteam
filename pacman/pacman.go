package pacman

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type direction int

const (
	none direction = iota
	right
	down
	left
	up
)

type Pacman struct {
	initX, initY          int
	sprite                [][]*ebiten.Image
	dir, nextDir, lastDir direction
	x, y                  int
	targetX, targetY      int
	game                  *Game
}

func (p *Pacman) draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	var spriteDirection int
	switch p.dir {
	case none:
		spriteDirection = int(p.lastDir) - 1
	default:
		spriteDirection = int(p.dir) - 1
	}
	op.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.sprite[spriteDirection][(p.x+p.y)/5%2], op)
}

func (p *Pacman) move() {
	if p.x == p.targetX && p.y == p.targetY {
		i := p.y / tileSize
		j := p.x / tileSize
		p.game.checkPill(i, j)
		p.nextTarget()
	}
	p.checkOpposites()
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

func (p *Pacman) checkOpposites() {
	if p.dir == right && p.nextDir == left {
		p.dir = p.nextDir
		p.targetX -= tileSize
	} else if p.dir == left && p.nextDir == right {
		p.dir = p.nextDir
		p.targetX += tileSize
	} else if p.dir == up && p.nextDir == down {
		p.dir = p.nextDir
		p.targetY += tileSize
	} else if p.dir == down && p.nextDir == up {
		p.dir = p.nextDir
		p.targetY -= tileSize
	}
}

func (p *Pacman) nextTarget() {

	if p.nextDir != none && !p.isWall(p.nextDir) {
		p.dir = p.nextDir
	} else if p.isWall(p.dir) {
		p.lastDir = p.dir
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

func (p *Pacman) isWall(dir direction) bool {

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

	if p.game.scene.stage[i][j] == wall {
		return true
	}
	return false
}

func (p *Pacman) getInput() {
	p.nextDir = none
	var keys []ebiten.Key
	keys = inpututil.AppendPressedKeys(keys)
	duration := math.MaxInt
	for _, key := range keys {
		if inpututil.KeyPressDuration(key) < duration {
			switch key {
			case ebiten.KeyArrowDown:
				p.nextDir = down
			case ebiten.KeyArrowUp:
				p.nextDir = up
			case ebiten.KeyArrowLeft:
				p.nextDir = left
			case ebiten.KeyArrowRight:
				p.nextDir = right
			default:
				continue
			}
			duration = inpututil.KeyPressDuration(key)
		}
	}
}

func (p *Pacman) reset() {
	p.x = p.initX
	p.y = p.initY
	p.targetX = p.initX
	p.targetY = p.initY
	p.dir = right
}
