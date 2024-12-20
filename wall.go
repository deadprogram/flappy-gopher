package main

import (
	rand "math/rand/v2"

	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	wallWidth        = 8   // 壁の幅
	wallHeight       = 128 // 壁の高さ
	holeYMax         = 72  // 穴のY座標の最大値
	holeHeight       = 64  // 穴のサイズ（高さ
	newWallsInterval = 200 // 新しい壁を追加する間隔
	wallStartX       = 240 // 壁の初期X座標
)

var (
	wallImage firefly.Image
)

func newWalls() *wallsData {
	wallImage = firefly.LoadFile("wall").Image()
	return &wallsData{
		walls: []*wallData{},
	}
}

type wallsData struct {
	walls []*wallData
}

func (w *wallsData) reset() {
	w.walls = []*wallData{}
}

func (w *wallsData) draw() {
	for _, wall := range w.walls {
		wall.draw()
	}
}

func (w *wallsData) score(l int) int {
	score := 0
	for i, wall := range w.walls {
		if wall.wallX < int(l) {
			score = i + 1
		}
	}
	return score
}

func (w *wallsData) move() {
	for _, wall := range w.walls {
		wall.move()
		if frames%4 == 0 {
			wall.move()
		}
	}
}

func (w *wallsData) add() {
	wall := &wallData{wallStartX, rand.N(holeYMax)}
	w.walls = append(w.walls, wall)
}

type wallData struct {
	wallX int
	holeY int
}

func (w *wallData) move() {
	w.wallX -= 1
}

func (w *wallData) draw() {
	// 上の壁の描画
	firefly.DrawImage(wallImage, firefly.Point{X: w.wallX, Y: w.holeY - wallHeight})

	// 下の壁の描画
	firefly.DrawImage(wallImage, firefly.Point{X: w.wallX, Y: w.holeY + holeHeight})
}

func (w *wallData) top() (int, int, int, int) {
	l := w.wallX
	t := w.holeY - wallHeight
	r := w.wallX + wallWidth
	b := w.holeY

	return l, t, r, b
}

func (w *wallData) bottom() (int, int, int, int) {
	l := w.wallX
	t := w.holeY + holeHeight
	r := w.wallX + wallWidth
	b := w.holeY + holeHeight + wallHeight

	return l, t, r, b
}
