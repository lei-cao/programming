package ui

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
)

var (
	barWidth      = 5
	barHeightUnit = 5
	barMargin     = 3
)

func NewRectSlice(values []int) *RectSlice {
	rs := new(RectSlice)
	rs.values = values
	rs.barImage, _ = ebiten.NewImage(rectSliceWidth(len(values)), rectSliceHeight(len(values)), ebiten.FilterDefault)
	rs.barImage.Fill(color.White)
	for k, v := range rs.values {
		r := NewRectangle(v)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(k*(barWidth+barMargin)), float64((len(rs.values)-r.V)*barHeightUnit))
		op.SourceRect = r.rect
		r.startOp = op
		r.startIndex = k
		r.endIndex = k

		rs.rectangles = append(rs.rectangles, r)
	}
	return rs
}

type RectSlice struct {
	values     []int
	rectangles []*Rectangle
	barImage   *ebiten.Image
}

func (rs *RectSlice) Update(progress float64) {
	for _, r := range rs.rectangles {
		if r.endIndex != r.startIndex {
			r.startOp.GeoM.Reset()
			dx := float64(r.startIndex*(barWidth+barMargin)) + progress*float64((r.endIndex-r.startIndex)*(barWidth+barMargin))
			r.startOp.GeoM.Translate(dx, float64((len(rs.values)-r.V)*barHeightUnit))
			if progress == 1 {
				r.startIndex = r.endIndex
			}
		}
	}
}

func (rs *RectSlice) NextStep(s visualizer.Stepper) {
	if step, ok := s.(*basicsort.Step); ok {
		if step.DoSwap() {
			rs.swap(step.A(), step.B())
		} else {
			rs.pass(step.A(), step.B())
		}
	}
}

func (rs *RectSlice) Draw(image *ebiten.Image) {
	for _, r := range rs.rectangles {
		image.DrawImage(rs.barImage, r.startOp)
	}
}

func (rs *RectSlice) swap(ia, ib int) {
	if ia < 0 || ib < 0 || ia > len(rs.rectangles)-1 || ib > len(rs.rectangles)-1 {
		return
	}

	ra := rs.rectangles[ia]
	rb := rs.rectangles[ib]
	ra.endIndex = ib
	rb.endIndex = ia
	rs.rectangles[ia] = rb
	rs.rectangles[ib] = ra
}

func (rs *RectSlice) pass(a, b int) {

}

func rectSliceWidth(size int) int {
	return size*barWidth + (size-1)*barMargin
}
func rectSliceHeight(size int) int {
	return size * barHeightUnit
}
