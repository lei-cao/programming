package ui

import (
	"image"
	"github.com/hajimehoshi/ebiten"
)

func NewRectangle(value int) *Rectangle {
	r := new(Rectangle)
	rect := image.Rect(0, 0, barWidth, value*barHeightUnit)
	r.rect = &rect
	r.V = value
	return r
}

type Rectangle struct {
	startOp   *ebiten.DrawImageOptions
	startIndex int
	endIndex int
	rect *image.Rectangle
	V    int
}
