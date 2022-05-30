package pacman

import (
	_ "image/png"
)

const (
	background_image_size = 100
	tile_size             = 32
)

type scene struct {
	stage *stage
}

func createScene(stage *stage) *scene {
	s := &scene{}
	s.stage = stage

	if s.stage == nil {
		s.stage = default_stage
	}

	return s
}
