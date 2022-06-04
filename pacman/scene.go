package pacman

import (
	_ "image/png"
)

const (
	backgroundImageSize = 100
	tileSize            = 32
)

type scene struct {
	stage          [][]tile
	pacmanInitialX int
	pacmanInitialY int
	enemyPositions [][]int
}

type tile int

const (
	empty tile = iota
	pill
	superPill
	wall
)

func createScene(level []string) *scene {
	s := &scene{}
	if level == nil {
		level = defaultLevel
	}

	s.stage = make([][]tile, len(level))

	for i, row := range level {
		s.stage[i] = make([]tile, len(row))
		for j, col := range row {
			switch col {
			case '#':
				s.stage[i][j] = wall
			case '.':
				s.stage[i][j] = pill
			case 'X':
				s.stage[i][j] = superPill
			case 'G':
				x := j * tileSize
				y := i * tileSize
				s.enemyPositions = append(s.enemyPositions, []int{x, y})
			case 'P':
				s.pacmanInitialX = j * tileSize
				s.pacmanInitialY = i * tileSize
			}
		}
	}

	return s
}
