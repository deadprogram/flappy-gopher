package main

import (
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

func updateGameover() {
	buttons = firefly.ReadButtons(firefly.Combined)
	if buttons.N || buttons.S || buttons.E || buttons.W {
		scene = gameStart
		score = 0
		frames = 0

		gopher.reset()
		walls.reset()
	}
}

func renderGameover() {
	firefly.ClearScreen(firefly.ColorBlue)

	gopher.draw()
	walls.draw()

	firefly.DrawText("GAME OVER", titleFont, firefly.Point{X: 90, Y: 60}, firefly.ColorWhite)
	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 90, Y: 100}, firefly.ColorWhite)
}
