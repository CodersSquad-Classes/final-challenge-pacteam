package pacman

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

var wall_img *ebiten.Image
var bg_img *ebiten.Image

func (s *scene) screenHeight() int {
	return len(s.stage.tile_matrix) * tile_size
}

func (s *scene) screenWidth() int {
	return len(s.stage.tile_matrix[0]) * tile_size
}

func createScene(stage *stage) *scene {
	s := &scene{}
	s.stage = stage

	if s.stage == nil {
		s.stage = default_stage
	}

	s.images = make(map[string]*ebiten.Image)

	s.loadImages()
	s.createStage()
	s.addTiles()

	return s
}

func (s *scene) loadImages() {
	var err_1 error
	var err_2 error

	wall_img, _, err_1 = ebitenutil.NewImageFromFile("assets/level.png")
	bg_img, _, err_2 = ebitenutil.NewImageFromFile("assets/background.png")

	if err_1 != nil {
		log.Fatal(err_1)
	}

	if err_2 != nil {
		log.Fatal(err_2)
	}

	s.images["wall"] = wall_img
	s.images["bg"] = bg_img
}

func (s *scene) createStage() {
	height := len(s.stage.tile_matrix)
	width := len(s.stage.tile_matrix[0])

	s.tile_matrix = make([][]string, height)

	for i := 0; i < height; i++ {
		s.tile_matrix[i] = make([]string, width)

		for j := 0; j < width; j++ {

		}
	}

}

func (s *scene) addTiles() {
	height := len(s.stage.tile_matrix)
	width := len(s.stage.tile_matrix[0])

	sizeW := ((width*tile_size)/background_image_size + 1) * background_image_size
	sizeH := ((height*tile_size)/background_image_size + 1) * background_image_size

	fmt.Println("sizeW: ", sizeW)
	fmt.Println("sizeH: ", sizeH)
	fmt.Println("width: ", width)
	fmt.Println("height: ", height)

	s.tile = ebiten.NewImage(sizeW, sizeH)

	// add background image
	for i := 0; i < sizeH/background_image_size; i++ {
		y := float64(i * background_image_size)

		for j := 0; j < sizeW/background_image_size; j++ {
			options := &ebiten.DrawImageOptions{}

			x := float64(j * background_image_size)

			options.GeoM.Translate(x, y)
			s.tile.DrawImage(bg_img, options)
		}
	}
}

func (s *scene) Update(screen *ebiten.Image) error {
	screen.Clear()
	screen.DrawImage(s.tile, nil)

	return nil
}
