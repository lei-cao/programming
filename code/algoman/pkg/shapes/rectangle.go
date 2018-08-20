package shapes

import (
	"image"
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/pkg/consts"
)

func NewRectangle(value int) *Rectangle {
	r := new(Rectangle)
	rect := image.Rect(0, 0, barWidth, value*barHeightUnit)
	r.barImage, _ = ebiten.NewImage(barWidth, value*barHeightUnit, ebiten.FilterDefault)
	r.barImage.Fill(consts.BarColor)
	r.rect = &rect
	r.V = value
	return r
}

type Rectangle struct {
	barImage   *ebiten.Image
	startOp    *ebiten.DrawImageOptions
	startIndex int
	endIndex   int
	isA        bool
	isB        bool
	rect       *image.Rectangle
	V          int
}
