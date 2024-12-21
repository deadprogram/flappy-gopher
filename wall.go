package main

import (
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

// wallData represents a single wall
type wallData struct {
	wallX int
	holeY int
}

// move moves the wall to the left
func (w *wallData) move() {
	w.wallX -= 1
}

// draw draws the wall
func (w *wallData) draw() {
	// draw the top wall
	firefly.DrawImage(wallImage, firefly.Point{X: w.wallX, Y: w.holeY - wallHeight})

	// draw the bottom wall
	firefly.DrawImage(wallImage, firefly.Point{X: w.wallX, Y: w.holeY + holeHeight})
}

// top returns the wall's top coordinates
func (w *wallData) top() (int, int, int, int) {
	l := w.wallX
	t := w.holeY - wallHeight
	r := w.wallX + wallWidth
	b := w.holeY

	return l, t, r, b
}

// bottom returns the wall's bottom coordinates
func (w *wallData) bottom() (int, int, int, int) {
	l := w.wallX
	t := w.holeY + holeHeight
	r := w.wallX + wallWidth
	b := w.holeY + holeHeight + wallHeight

	return l, t, r, b
}
