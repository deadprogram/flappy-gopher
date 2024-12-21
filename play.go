package main

import (
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

// updatePlay updates the game play
func updatePlay() {
	frames += 1

	buttons = firefly.ReadButtons(firefly.GetMe())
	if buttons.N || buttons.S || buttons.E || buttons.W {
		gopher.jump()
	}

	// add additional walls?
	if frames%newWallsInterval == 0 {
		walls.add()
	}

	// current score
	score = gopher.score(walls)

	gopher.move()
	walls.move()

	switch {
	case gopher.hitWalls(walls):
		scene = gameOver
	case gopher.hitBottom():
		scene = gameOver
	case gopher.hitTop():
		scene = gameOver
	}
}

// renderPlay renders the game play onto the screen
func renderPlay() {
	firefly.ClearScreen(firefly.ColorWhite)

	gopher.draw()
	walls.draw()

	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 10, Y: 10}, firefly.ColorDarkBlue)
}
