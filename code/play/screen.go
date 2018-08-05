package controller

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
)

var barWidth = 8
var barSpace = 2
var heightUnit = 5

// The screen including the elements on the canvas
// Maintain the ready for next state
type Screen struct {
	Ctx            *canvas.Context2D
	Rectangles      []*Rectangle
	FinishedDrawing map[int]bool
	Ready           bool
	AIndex          int
	BIndex          int
}

func (s *Screen) Draw(delta float64) {
	s.draw(delta)
}

func (s *Screen) draw(delta float64) {
	for k, r := range s.Rectangles {
		r.IsB = false
		r.IsA = false
		if s.AIndex == r.Index {
			r.IsA = true
		}
		if s.BIndex == r.Index {
			r.IsB = true
		}
		s.FinishedDrawing[k] = r.Update(delta)
		s.drawRect(r)
	}
	for _, finished := range s.FinishedDrawing {
		if !finished {
			s.Ready = false
			return
		}
	}
	s.Ready = true
}

func (s *Screen) drawRect(r *Rectangle) {
	s.Ctx.FillStyle = "#B9314F"
	if r.IsA {
		s.Ctx.FillStyle = "#2E86AB"
	} else if r.IsB {
		//ctx.FillStyle = "#12355B"
		s.Ctx.FillStyle = "green"
	}
	s.Ctx.FillRect(r.Left, r.Top, r.Width, r.Height)
}

func (s *Screen) swap(ia, ib int) {
	a := s.Rectangles[ia]
	b := s.Rectangles[ib]
	a.IsA = true
	b.IsB = true
	s.AIndex = ia
	s.BIndex = ib

	a.ToIndex = ib
	b.ToIndex = ia

	s.Rectangles[ia] = b
	s.Rectangles[ib] = a
}

func (s *Screen) pass(ia, ib int) {
	a := s.Rectangles[ia]
	b := s.Rectangles[ib]
	a.IsA = true
	b.IsB = true
	a.Waiting = 2
	s.AIndex = ia
	s.BIndex = ib
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
