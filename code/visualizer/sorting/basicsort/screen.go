package basicsort

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/visualizer"
	"github.com/lei-cao/programming/code/visualizer/defaults"
	basic "github.com/lei-cao/programming/code/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/visualizer/ui"
)

func NewScreen(id string, size int, nums []int) *Screen {
	s := new(Screen)
	s.id = id
	s.size = size

	obj := createCanvas(s.id, s.size)
	s.c = canvas.New(obj)
	s.ctx = s.c.GetContext2D()

	pa := ui.Point{X: 0, Y: 0}
	s.rs = ui.NewRectSlice(s.ctx, nums, pa, "a", false)

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
	rs              *ui.RectSlice
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

// TODO type assertion. Have to?
func (s *Screen) Update(i visualizer.Stepper) {
	if step, ok := i.(*basic.Step); ok {
		if step.DoSwap() {
			s.Swap(step.A(), step.B())
		} else {
			s.Pass(step.A(), step.B())
		}
	}
}

func (s *Screen) Swap(ia, ib int) {
	s.rs.Swap(ia, ib)
}

func (s *Screen) Pass(ia, ib int) {
	s.rs.Pass(ia, ib)
}

func (s *Screen) draw(progress float64) {
	s.ctx.FillStyle = defaults.DefaultColor.BackgroundColor
	s.ctx.FillRect(0, 0, float64(s.c.Width), float64(s.c.Height))

	s.rs.Draw(progress)

	s.ready = s.rs.Ready()
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
	return defaults.BarWidth*size + (size-1)*defaults.BarSpace
}

func canvasHeight(size int) int {
	return size * defaults.HeightUnit
}
