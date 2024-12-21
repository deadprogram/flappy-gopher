package main

import "github.com/firefly-zero/firefly-go/firefly"

// logic for the start scene
func updateStart() {
	frames += 1
	if frames > 60 {
		buttons = firefly.ReadButtons(firefly.GetMe())
		if buttons.N || buttons.S || buttons.E || buttons.W {
			frames = 0
			scene = gamePlay
		}
	}
}

// render the start scene
func renderStart() {
	firefly.ClearScreen(firefly.ColorBlue)
	firefly.DrawText("FLAPPY GOPHER", titleFont, firefly.Point{X: 80, Y: 60}, firefly.ColorWhite)
	firefly.DrawText("Press any button to Start", titleFont, firefly.Point{X: 44, Y: 100}, firefly.ColorWhite)
}
