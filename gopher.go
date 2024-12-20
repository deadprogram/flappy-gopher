package main

import "github.com/firefly-zero/firefly-go/firefly"

const (
	gopherWidth  = 24
	gopherHeight = 24

	g    = 0.05 // Gravity (重力加速度) の略
	jump = -1.0 // ジャンプ力
)

type gopherData struct {
	img firefly.Image
	x   float32
	y   float32
	vy  float32 // Velocity of y (速度のy成分) の略
}

func newGopher() *gopherData {
	return &gopherData{
		img: firefly.LoadFile("gopher").Image(),
		x:   20,
		y:   100,
		vy:  0,
	}
}

func (gd *gopherData) reset() {
	gd.x = 20
	gd.y = 100
	gd.vy = 0
}

func (gd *gopherData) jump() {
	gd.vy = jump
}

func (gd *gopherData) move() {
	gd.vy += g    // 速度に加速度を足す
	gd.y += gd.vy // 位置に速度を足す
}

func (gd *gopherData) draw() {
	firefly.DrawImage(gd.img, firefly.Point{X: int(gd.x), Y: int(gd.y)})
}

func (gd *gopherData) position() (int, int, int, int) {
	l := int(gd.x)
	t := int(gd.y)
	r := int(gd.x) + gopherWidth
	b := int(gd.y) + gopherHeight

	return l, t, r, b
}

func (gd *gopherData) isHit(wl, wt, wr, wb int) bool {
	gl, gt, gr, gb := gd.position()

	if gr < wl || gl > wr {
		return false
	}

	if gb < wt || gt > wb {
		return false
	}

	return true
}

func (gd *gopherData) hitWalls(w *wallsData) bool {
	for _, wall := range w.walls {
		// 上の壁を表す四角形を作る
		bLeft, bTop, bRight, bBottom := wall.top()

		// 上の壁との当たり判定
		if gd.isHit(bLeft, bTop, bRight, bBottom) {
			return true
		}

		// 下の壁を表す四角形を作る
		bLeft, bTop, bRight, bBottom = wall.bottom()

		// 下の壁との当たり判定
		if gopher.isHit(bLeft, bTop, bRight, bBottom) {
			return true
		}
	}

	return false
}

func (gd *gopherData) hitTop() bool {
	return gd.y < 0
}

func (gd *gopherData) hitBottom() bool {
	b := int(gd.y) + gopherHeight
	return b > 160
}

func (gd *gopherData) score(w *wallsData) int {
	l, _, _, _ := gd.position()
	return w.score(l)
}
