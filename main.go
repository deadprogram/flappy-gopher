package main

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	gameStart = "start"
	gamePlay  = "game"
	gameOver  = "gameover"
)

var (
	frames = 0 // 経過フレーム数
	score  = 0
	scene  = gameStart

	titleFont firefly.Font
	gopher    *gopherData
	walls     *wallsData

	buttons firefly.Buttons
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

func boot() {
	titleFont = firefly.LoadFile("font", nil).Font()
	gopher = newGopher()
	walls = newWalls()
}

func update() {
	switch scene {
	case gameStart:
		updateStart()
	case gamePlay:
		updatePlay()
	case gameOver:
		updateGameover()
	}
}

func render() {
	switch scene {
	case gameStart:
		renderStart()
	case gamePlay:
		renderPlay()
	case gameOver:
		renderGameover()
	}
}

func main() {
}
