package main

import "github.com/firefly-zero/firefly-go/firefly"

const (
	wallWidth  = 8   // 壁の幅
	wallHeight = 128 // 壁の高さ
	holeYMax   = 48  // 穴のY座標の最大値
	holeHeight = 40  // 穴のサイズ（高さ）
)

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
