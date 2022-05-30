package pacman

import (
	_ "image/png"
)

const (
	backgroundImageSize = 100
	tileSize            = 32
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
