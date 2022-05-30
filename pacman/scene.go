package pacman

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	background_image_size = 100
	tile_size             = 32
	tile_x_num            = 40
)

type scene struct {
	tile_matrix [][]string
	tile        *ebiten.Image
	images      map[string]*ebiten.Image
	stage       *stage
}

func createScene(stage *stage) *scene {
	s := &scene{}
	s.stage = stage

	if s.stage == nil {
		s.stage = default_stage
	}

	return s
}
