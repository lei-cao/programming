package mergesort

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/learning-cs-again/code/visualizer/defaults"
)

func NewRectSlice(ctx *canvas.Context2D, nums []int, startPoint *Point, name string, displayName bool) *RectSlice {
	rs := &RectSlice{}
	rs.nums = nums
	rs.StartPoint = startPoint
	rs.Size = len(nums)
	for k, v := range nums {
		r := NewRect(ctx, rs.rectPoint(k, v), defaults.BarWidth, defaults.HeightUnit*float64(v), v)
		rs.AddRect(r)
	}
	rs.Name = name
	rs.DisplayName = displayName
	return rs
}

func rectSliceHeight(size int) float64 {
	return defaults.HeightUnit * float64(size)
}

type RectSlice struct {
	Rectangles  []*Rectangle
	StartPoint  *Point
	Size        int
	Name        string
	DisplayName bool
	nums        []int
}

func (rs *RectSlice) AddRect(rectangle *Rectangle) {
	rs.Rectangles = append(rs.Rectangles, rectangle)
}

func (rs *RectSlice) Height() float64 {
	return float64(rs.Size) * defaults.HeightUnit
}

func (rs *RectSlice) rectPoint(k int, v int) *Point {
	p := &Point{
		rs.StartPoint.X + (defaults.BarWidth+defaults.BarSpace)*float64(k),
		rs.StartPoint.Y + rectSliceHeight(rs.Size) - float64(v)*defaults.HeightUnit,
	}
	return p
}

func NewRect(ctx *canvas.Context2D, startPoint *Point, width float64, height float64, value int) *Rectangle {
	r := new(Rectangle)
	r.Ctx = ctx
	r.startPoint = startPoint
	r.destPoint = startPoint
	r.Width = width
	r.Height = height
	r.V = value
	r.Color = defaults.DefaultColor.BarColor
	return r
}

// Represent the element in the problem slice
type Rectangle struct {
	Ctx        *canvas.Context2D
	startPoint *Point
	destPoint  *Point
	Width      float64
	Height     float64
	Color      string
	V          int
}

func (r *Rectangle) Animate(progress float64) bool {
	finished := r.update(progress)
	r.draw()
	return finished
}

func (r *Rectangle) update(progress float64) bool {
	var finished bool

	r.startPoint.MoveTo(r.destPoint, progress)

	if progress == 1 {
		r.startPoint = r.destPoint
		finished = true
	}
	return finished
}

func (r *Rectangle) draw() {
	r.Ctx.FillStyle = r.Color
	if r.moving() {
		r.Ctx.FillStyle = defaults.DefaultColor.AColor
	}
	r.Ctx.FillRect(r.startPoint.X, r.startPoint.Y, r.Width, r.Height)
}

func (r *Rectangle) moving() bool {
	return !r.startPoint.Equals(r.destPoint)
}
