package main

import (
	"math/rand"
)

type Enemy struct {
	xPos  int
	yPos  int
	color [4]float64
}

func (e *Enemy) moveRandom() {
	dir := rand.Intn(3)
	switch dir {
	case 0: //up
		if isWall(e.xPos, e.yPos-1) {
			e.yPos -= 1
		} else {
			e.moveRandom()
		}
	case 1: //down
		if isWall(e.xPos, e.yPos+1) {
			e.yPos += 1
		} else {
			e.moveRandom()
		}
	case 2: //right
		if isWall(e.xPos+1, e.yPos) {
			e.xPos += 1
		} else {
			e.moveRandom()
		}
	default: //left
		if isWall(e.xPos-1, e.yPos) {
			e.xPos -= 1
		} else {
			e.moveRandom()
		}
	}
}
func isWall(x int, y int) bool {
	return true
}
