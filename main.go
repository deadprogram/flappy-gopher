package main

import (
	rand "math/rand/v2"
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

var (
	frames           = 0             // 経過フレーム数
	newWallsInterval = 200           // 新しい壁を追加する間隔
	wallStartX       = 240           // 壁の初期X座標
	walls            = []*wallData{} // 壁のX座標とY座標

	scene = "title"
	score = 0

	titleFont firefly.Font
	gopher    *gopherData
	wallImage firefly.Image

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
	wallImage = firefly.LoadFile("wall").Image()
}

func update() {
	switch scene {
	case "title":
		buttons = firefly.ReadButtons(firefly.GetMe())
		if buttons.N || buttons.S || buttons.E || buttons.W {
			scene = "game"
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
			walls = []*wallData{}
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
	firefly.DrawText("Press any button to Start", titleFont, firefly.Point{X: 100, Y: 100}, firefly.ColorWhite)
}

func renderGame() {
	firefly.ClearScreen(firefly.ColorWhite)
	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 10, Y: 10}, firefly.ColorDarkBlue)

	gopher.draw()

	for _, wall := range walls {
		wall.draw()
	}
}

func renderGameover() {
	firefly.ClearScreen(firefly.ColorBlue)
	gopher.draw()

	for _, wall := range walls {
		wall.draw()
	}

	firefly.DrawText("GAME OVER", titleFont, firefly.Point{X: 100, Y: 100}, firefly.ColorWhite)
	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 100, Y: 120}, firefly.ColorWhite)
}

func updateGame() {
	buttons = firefly.ReadButtons(firefly.GetMe())
	if buttons.N || buttons.S || buttons.E || buttons.W {
		gopher.jump()
	}

	frames += 1

	// add walls
	if frames%newWallsInterval == 0 {
		wall := &wallData{wallStartX, rand.N(holeYMax)}
		walls = append(walls, wall)
	}

	// current score
	l, _, _, _ := gopher.position()
	for i, wall := range walls {
		if wall.wallX < int(l) {
			score = i + 1
		}
	}

	for _, wall := range walls {
		wall.move()
		if frames%4 == 0 {
			wall.move()
		}
	}

	gopher.move()

	for _, wall := range walls {
		// 上の壁を表す四角形を作る
		bLeft, bTop, bRight, bBottom := wall.top()

		// 上の壁との当たり判定
		if gopher.isHit(bLeft, bTop, bRight, bBottom) {
			scene = "gameover"
		}

		// 下の壁を表す四角形を作る
		bLeft, bTop, bRight, bBottom = wall.bottom()

		// 下の壁との当たり判定
		if gopher.isHit(bLeft, bTop, bRight, bBottom) {
			scene = "gameover"
		}

		_, t, _, b := gopher.position()
		if t < 0 {
			scene = "gameover"
		}
		if b > 160 {
			scene = "gameover"
		}
	}
}

func main() {
}
