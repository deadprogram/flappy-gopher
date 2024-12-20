package main

import (
	rand "math/rand/v2"
	"strconv"

	"github.com/firefly-zero/firefly-go/firefly"
)

type wallData struct {
	wallX int
	holeY int
}

var (
	x    = 20.0
	y    = 30.0
	vy   = 0.0  // Velocity of y (速度のy成分) の略
	g    = 0.25 // Gravity (重力加速度) の略
	jump = -1.0 // ジャンプ力

	frames               = 0             // 経過フレーム数
	moveInterval         = 5             // 壁の追加間隔
	newWallsInterval     = 200           // 新しい壁を追加する間隔
	wallMovementInterval = 10            // 壁の移動量
	wallStartX           = 240           // 壁の初期X座標
	walls                = []*wallData{} // 壁のX座標とY座標
	wallWidth            = 8             // 壁の幅
	wallHeight           = 128           // 壁の高さ
	holeYMax             = 48            // 穴のY座標の最大値
	holeHeight           = 40            // 穴のサイズ（高さ）

	gopherWidth  = 24
	gopherHeight = 24

	scene = "title"
	score = 0

	titleFont firefly.Font
	gopher    firefly.Image
	wall      firefly.Image

	buttons firefly.Buttons
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

func boot() {
	titleFont = firefly.LoadFile("font").Font()
	gopher = firefly.LoadFile("gopher").Image()
	wall = firefly.LoadFile("wall").Image()
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
			x = 20.0
			y = 30.0
			vy = 0.0
			frames = 0
			walls = []*wallData{}
			score = 0
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

	firefly.DrawImage(gopher, firefly.Point{X: int(x), Y: int(y)})

	for _, wall := range walls {
		drawWalls(wall)
	}
}

func renderGameover() {
	firefly.ClearScreen(firefly.ColorBlue)
	firefly.DrawImage(gopher, firefly.Point{X: int(x), Y: int(y)})

	for _, wall := range walls {
		drawWalls(wall)
	}

	firefly.DrawText("GAME OVER", titleFont, firefly.Point{X: 100, Y: 100}, firefly.ColorWhite)
	firefly.DrawText("Score: "+strconv.Itoa(score), titleFont, firefly.Point{X: 100, Y: 120}, firefly.ColorWhite)
}

func drawWalls(w *wallData) {
	// 上の壁の描画
	firefly.DrawImage(wall, firefly.Point{X: w.wallX, Y: w.holeY - wallHeight})

	// 下の壁の描画
	firefly.DrawImage(wall, firefly.Point{X: w.wallX, Y: w.holeY + holeHeight})
}

func updateGame() {
	buttons = firefly.ReadButtons(firefly.GetMe())
	if buttons.N || buttons.S || buttons.E || buttons.W {
		vy = jump
	}

	frames += 1

	// add walls
	if frames%newWallsInterval == 0 {
		wall := &wallData{wallStartX, rand.N(holeYMax)}
		walls = append(walls, wall)
	}

	// current score
	for i, wall := range walls {
		if wall.wallX < int(x) {
			score = i + 1
		}
	}

	if frames%wallMovementInterval == 0 {
		for _, wall := range walls {
			wall.wallX -= 8 // 少しずつ左へ
		}
	}

	if frames%moveInterval == 0 {
		vy += g // 速度に加速度を足す
		y += vy // 位置に速度を足す

		for _, wall := range walls {
			// gopherくんを表す四角形を作る
			aLeft := int(x)
			aTop := int(y)
			aRight := int(x) + gopherWidth
			aBottom := int(y) + gopherHeight

			// 上の壁を表す四角形を作る
			bLeft := wall.wallX
			bTop := wall.holeY - wallHeight
			bRight := wall.wallX + wallWidth
			bBottom := wall.holeY

			// 上の壁との当たり判定
			if hitTestRects(aLeft, aTop, aRight, aBottom, bLeft, bTop, bRight, bBottom) {
				scene = "gameover"
			}

			// 下の壁を表す四角形を作る
			bLeft = wall.wallX
			bTop = wall.holeY + holeHeight
			bRight = wall.wallX + wallWidth
			bBottom = wall.holeY + holeHeight + wallHeight

			// 下の壁との当たり判定
			if hitTestRects(aLeft, aTop, aRight, aBottom, bLeft, bTop, bRight, bBottom) {
				scene = "gameover"
			}

			if y < 0 {
				scene = "gameover"
			}
			if y > 160 {
				scene = "gameover"
			}
		}
	}
}

func hitTestRects(aLeft, aTop, aRight, aBottom, bLeft, bTop, bRight, bBottom int) bool {
	return aLeft < bRight &&
		bLeft < aRight &&
		aTop < bBottom &&
		bTop < aBottom
}

func main() {
}
