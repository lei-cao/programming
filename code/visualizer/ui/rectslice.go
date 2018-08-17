package ui

import (
	"github.com/lei-cao/programming/code/visualizer/defaults"
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/visualizer"
	basic "github.com/lei-cao/programming/code/algorithms/sorting/basicsort"
)

func NewRectSlice(ctx *canvas.Context2D, nums []int, startPoint Point, name string, displayName bool) *RectSlice {
	rs := &RectSlice{}
	rs.Element = NewElement()
	rs.StartPoint = startPoint
	rs.Size = len(nums)
	for k, v := range nums {
		r := NewRect(ctx, rs.RectPoint(k, v), defaults.BarWidth, defaults.HeightUnit*float64(v), k, v)
		rs.AddRect(r)
	}
	rs.finishedDrawing = make(map[int]bool)
	rs.Id = name
	rs.aIndex = -1
	rs.bIndex = -1
	return rs
}

func RectSliceHeight(size int) float64 {
	return defaults.HeightUnit * float64(size)
}

func RectSliceWidth(size int) float64 {
	return float64(size)*defaults.BarWidth + float64(size-1)*defaults.BarSpace
}

type RectSlice struct {
	*Element
	Rectangles      []*Rectangle
	StartPoint      Point
	Size            int
	finishedDrawing map[int]bool
	aIndex          int
	bIndex          int
}

func (rs *RectSlice) Update(stepper visualizer.Stepper) {
	if step, ok := stepper.(*basic.Step); ok {
		if step.DoSwap() {
			rs.Swap(step.A(), step.B())
		} else {
			rs.Pass(step.A(), step.B())
		}
	}
}

func (rs *RectSlice) Draw(progress float64) {
	for k, r := range rs.Rectangles {
		r.isB = false
		r.isA = false
		if rs.aIndex == r.Index {
			r.isA = true
		}
		if rs.bIndex == r.Index {
			r.isB = true
		}
		rs.finishedDrawing[k] = r.Animate(progress)
	}

}

func (rs *RectSlice) Ready() bool {
	for _, finished := range rs.finishedDrawing {
		if !finished {
			return false
		}
	}
	return true
}

func (rs *RectSlice) AddRect(rectangle *Rectangle) {
	rs.Rectangles = append(rs.Rectangles, rectangle)
}

func (rs *RectSlice) Width() float64 {
	return RectSliceWidth(rs.Size)
}

func (rs *RectSlice) Height() float64 {
	return RectSliceHeight(rs.Size)
}

func (rs *RectSlice) RectPoint(k int, v int) Point {
	p := Point{
		rs.StartPoint.X + (defaults.BarWidth+defaults.BarSpace)*float64(k),
		rs.StartPoint.Y + RectSliceHeight(rs.Size) - float64(v)*defaults.HeightUnit,
	}
	return p
}

func (rs *RectSlice) Swap(ia, ib int) {
	if ia < 0 || ib < 0 || ia > len(rs.Rectangles)-1 || ib > len(rs.Rectangles)-1 {
		return
	}

	ra := rs.Rectangles[ia]
	rb := rs.Rectangles[ib]

	rs.aIndex = ia
	rs.bIndex = ib
	ra.ToIndex = ib
	rb.ToIndex = ia

	ra.isA = true
	rb.isB = true

	ra.DestPoint = rs.RectPoint(ib, ra.V)
	rb.DestPoint = rs.RectPoint(ia, rb.V)

	rs.Rectangles[ia] = rb
	rs.Rectangles[ib] = ra
}

func (rs *RectSlice) Pass(ia, ib int) {
	if ia < 0 || ib < 0 || ia > len(rs.Rectangles)-1 || ib > len(rs.Rectangles)-1 {
		return
	}
	rs.Rectangles[ia].isA = true
	rs.Rectangles[ib].isB = true

	rs.aIndex = ia
	rs.bIndex = ib
}
