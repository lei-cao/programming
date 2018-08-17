package basicsort

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/v2/visualizer/ui"
	"github.com/lei-cao/programming/code/v2/visualizer/sorting"
)

func NewScreen(id string, size int, nums []int) *Screen {
	s := new(Screen)
	s.Id = id
	s.Size = size

	obj := createCanvas(s.Id, s.Size, ui.RectSliceWidth(size), ui.RectSliceHeight(size))
	s.C = canvas.New(obj)
	s.Ctx = s.C.GetContext2D()

	pa := ui.Point{X: 0, Y: 0}
	s.Element = ui.NewRectSlice(s.Ctx, nums, pa, "a", false)

	return s
}

// The screen including the elements on the canvas
// Maintain the ready for next state
type Screen struct {
	sorting.BaseScreen
}

func createCanvas(id string, size int, width float64, height float64) *js.Object {
	body := js.Global.Get("document").Call("getElementById", id)
	obj := js.Global.Get("document").Call("createElement", "canvas")
	obj.Set("width", strconv.Itoa(int(width)))
	obj.Set("height", strconv.Itoa(int(height)))
	body.Set("innerHTML", "")
	body.Call("appendChild", obj)
	return obj
}
