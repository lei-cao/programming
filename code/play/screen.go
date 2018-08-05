package controller

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
)

// The screen including the elements on the canvas
// Maintain the ready for next state
type Screen struct {
	Rectangles      []*Rectangle
	FinishedDrawing map[int]bool
	Ready           bool
	AIndex          int
	BIndex          int
}

func (s *Screen) draw(ctx *canvas.Context2D, delta float64) {
	for k, r := range s.Rectangles {
		r.IsB = false
		r.IsA = false
		if s.AIndex == r.Index {
			r.IsA = true
		}
		if s.BIndex == r.Index {
			r.IsB = true
		}
		s.FinishedDrawing[k] = r.Draw(ctx, delta)
	}
	for _, finished := range s.FinishedDrawing {
		if !finished {
			s.Ready = false
			return
		}
	}
	s.Ready = true
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
