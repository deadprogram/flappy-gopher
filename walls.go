package main

import (
	"math/rand/v2"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	wallImage firefly.Image
)

// wallData represents all of the walls in the current game
type wallsData struct {
	walls []*wallData
}

func newWalls() *wallsData {
	wallImage = firefly.LoadFile("wall", nil).Image()
	return &wallsData{
		walls: []*wallData{},
	}
}

// reset clears all walls
func (w *wallsData) reset() {
	w.walls = []*wallData{}
}

// draw draws all walls
func (w *wallsData) draw() {
	for _, wall := range w.walls {
		wall.draw()
	}
}

// score returns the current score of the player based on the gopher's position
func (w *wallsData) score(l int) int {
	score := 0
	for i, wall := range w.walls {
		if wall.wallX < int(l) {
			score = i + 1
		}
	}
	return score
}

// move moves all walls
func (w *wallsData) move() {
	for _, wall := range w.walls {
		wall.move()
		if frames%4 == 0 {
			wall.move()
		}
	}
}

// hitWalls returns true if the gopher hits any walls
func (w *wallsData) add() {
	wall := &wallData{wallStartX, rand.N(holeYMax)}
	w.walls = append(w.walls, wall)
}
