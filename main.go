package main

import (
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	frames = 0 // 経過フレーム数
	scene  = "title"
	score  = 0

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
	titleFont = firefly.LoadFile("font").Font()
	gopher = newGopher()
	walls = newWalls()
}

func update() {
	switch scene {
	case "title":
		frames += 1
		if frames > 60 {
			buttons = firefly.ReadButtons(firefly.GetMe())
			if buttons.N || buttons.S || buttons.E || buttons.W {
				scene = "game"
				frames = 0
			}
		}
	case "game":
		updateGame()

	case "gameover":
		buttons = firefly.ReadButtons(firefly.GetMe())
		if buttons.N || buttons.S || buttons.E || buttons.W {
			scene = "title"
			score = 0
			frames = 0

			gopher.reset()
			walls.reset()
		}
	}
}

func render() {
	switch scene {
	case "title":
		renderTitle()
	case "game":
		renderGame()
	case "gameover":
		renderGameover()
	}
}

func renderTitle() {
	firefly.ClearScreen(firefly.ColorBlue)
	firefly.DrawText("FLAPPY GOPHER", titleFont, firefly.Point{X: 80, Y: 60}, firefly.ColorWhite)
	firefly.DrawText("Press any button to Start", titleFont, firefly.Point{X: 44, Y: 100}, firefly.ColorWhite)
}

func renderGame() {
	firefly.ClearScreen(firefly.ColorWhite)

	gopher.draw()
	walls.draw()

	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 10, Y: 10}, firefly.ColorDarkBlue)
}

func renderGameover() {
	firefly.ClearScreen(firefly.ColorBlue)

	gopher.draw()
	walls.draw()

	firefly.DrawText("GAME OVER", titleFont, firefly.Point{X: 90, Y: 60}, firefly.ColorWhite)
	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 90, Y: 100}, firefly.ColorWhite)
}

func updateGame() {
	buttons = firefly.ReadButtons(firefly.GetMe())
	if buttons.N || buttons.S || buttons.E || buttons.W {
		gopher.jump()
	}

	frames += 1

	// add walls
	if frames%newWallsInterval == 0 {
		walls.add()
	}

	// current score
	score = gopher.score(walls)

	gopher.move()
	walls.move()

	switch {
	case gopher.hitWalls(walls):
		scene = "gameover"
	case gopher.hitBottom():
		scene = "gameover"
	case gopher.hitTop():
		scene = "gameover"
	}
}

func main() {
}
