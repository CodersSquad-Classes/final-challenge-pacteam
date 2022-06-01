package pacman

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	xPos  int
	yPos  int
	color [4]float64
}

func (e *Enemy) moveRandom() {
	e.xPos += 30
	e.yPos += 30
	fmt.Print(" x: ", e.xPos, " y: ", e.yPos)
}
func (e *Enemy) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	e.moveRandom()
	options.GeoM.Translate(float64(e.xPos), float64(e.yPos))
	options.ColorM.Translate(e.color[0], e.color[1], e.color[2], e.color[3])
	screen.DrawImage(ghost, options)

}

/*
func (e *Enemy) moveRandom() {
	dir := rand.Intn(3)
	switch dir {
	case 0: //up
		if isWall(e.xPos, e.yPos-10) {
			e.moveRandom()
		} else {
			e.yPos -= 10
		}
	case 1: //down
		if isWall(e.xPos, e.yPos+10) {
			e.moveRandom()
		} else {
			e.yPos += 10
		}
	case 2: //right
		if isWall(e.xPos+10, e.yPos) {
			e.moveRandom()
		} else {
			e.xPos += 10
		}
	default: //left
		if isWall(e.xPos-10, e.yPos) {
			e.moveRandom()
		} else {
			e.xPos -= 10
		}
	}
}*/
func isWall(x int, y int) bool {
	return true
}
