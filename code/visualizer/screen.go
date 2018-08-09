package visualizer

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
)

var barWidth = 8
var barSpace = 2
var heightUnit = 5

func NewScreen(id string, size int, nums []int) *Screen {
	s := new(Screen)
	s.id = id
	s.size = size

	obj := createCanvas(s.id, s.size)
	s.c = canvas.New(obj)
	s.ctx = s.c.GetContext2D()

	s.rectangles = []*Rectangle{}

	for k, v := range nums {
		r := NewRect(size, k, v, s.ctx)
		s.rectangles = append(s.rectangles, r)
	}
	s.finishedDrawing = map[int]bool{}
	return s
}

// The screen including the elements on the canvas
// Maintain the ready for next state
type Screen struct {
	id              string
	size            int
	c               *canvas.Canvas
	ctx             *canvas.Context2D
	rectangles      []*Rectangle
	finishedDrawing map[int]bool
	ready           bool
	aIndex          int
	bIndex          int
}

func (s *Screen) Ready() bool {
	return s.ready
}

func (s *Screen) Clear() {
	s.ctx.ClearRect(0, 0, float64(canvasWidth(s.size)), float64(canvasHeight(s.size)))
}

func (s *Screen) Draw(progress float64) {
	s.draw(progress)
}

func (s *Screen) Update(i Stepper) {
	if i.DoSwap() {
		s.Swap(i.A(), i.B())
	} else {
		s.Pass(i.A(), i.B())
	}
}

func (s *Screen) Swap(ia, ib int) {
	a := s.rectangles[ia]
	b := s.rectangles[ib]
	a.IsA = true
	b.IsB = true
	s.aIndex = ia
	s.bIndex = ib

	a.ToIndex = ib
	b.ToIndex = ia

	s.rectangles[ia] = b
	s.rectangles[ib] = a
}

func (s *Screen) Pass(ia, ib int) {
	a := s.rectangles[ia]
	b := s.rectangles[ib]
	a.IsA = true
	b.IsB = true
	s.aIndex = ia
	s.bIndex = ib
}

func (s *Screen) draw(progress float64) {
	s.ctx.FillStyle = defaultColor.BackgroundColor
	s.ctx.FillRect(0, 0, float64(s.c.Width), float64(s.c.Height))
	for k, r := range s.rectangles {
		r.IsB = false
		r.IsA = false
		if s.aIndex == r.Index {
			r.IsA = true
		}
		if s.bIndex == r.Index {
			r.IsB = true
		}
		s.finishedDrawing[k] = r.Animate(progress)
	}
	for _, finished := range s.finishedDrawing {
		if !finished {
			s.ready = false
			return
		}
	}
	s.ready = true
}

func createCanvas(id string, size int) *js.Object {
	body := js.Global.Get("document").Call("getElementById", id)
	obj := js.Global.Get("document").Call("createElement", "canvas")
	obj.Set("width", strconv.Itoa(canvasWidth(size)))
	obj.Set("height", strconv.Itoa(canvasHeight(size)))
	body.Set("innerHTML", "")
	body.Call("appendChild", obj)
	return obj
}

func canvasWidth(size int) int {
	return barWidth*size + (size-1)*barSpace
}

func canvasHeight(size int) int {
	return size * heightUnit
}
