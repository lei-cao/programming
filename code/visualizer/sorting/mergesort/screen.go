package mergesort

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/visualizer"
	"github.com/lei-cao/programming/code/algorithms/sorting/mergesort"
	"github.com/lei-cao/programming/code/visualizer/defaults"
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
	s.rsA = ui.NewRectSlice(s.ctx, nums, pa, "a", false)

	pb := ui.Point{X: 0, Y: s.rsA.Height() + 30}
	s.rsB = ui.NewRectSlice(s.ctx, nums, pb, "b", false)

	s.aName = "a"
	s.bName = "b"
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
	rsA             *ui.RectSlice
	rsB             *ui.RectSlice
	aName           string
	bName           string
	finishedDrawing map[int]bool
	ready           bool
	iBegin          int
	iMid            int
	iEnd            int
}

func (s *Screen) Ready() bool {
	return s.ready
}

func (s *Screen) Clear() {
	s.ctx.ClearRect(0, 0, float64(canvasWidth(s.size)), float64(canvasHeight(s.size, 2)))
}

func (s *Screen) Draw(progress float64) {
	s.draw(progress)
}

func (s *Screen) Update(i visualizer.Stepper) {
	if step, ok := i.(*mergesort.Step); ok {
		if step.IsFirstStep() {
		}

		if step.IsSplitStep() {
			// split a step
			s.splitA(step)
			s.splitB(step)
		}

		if step.IsAssignStep() {
			s.assign(step)
		}
	}
}

func (s *Screen) assign(step *mergesort.Step) {
	var rsFrom *ui.RectSlice
	var rsTo *ui.RectSlice
	if s.aName == step.From {
		rsFrom = s.rsA
		rsTo = s.rsB
	} else {
		rsFrom = s.rsB
		rsTo = s.rsA
	}

	rFrom := rsFrom.Rectangles
	rTo := rsTo.Rectangles

	var r *ui.Rectangle
	var i int
	if step.Assign == "i" {
		i = step.I
		r = rFrom[step.I]
		rFrom[step.I] = r
	} else {
		i = step.J
		r = rFrom[step.J]
		rFrom[step.J] = r
	}
	dest := rsTo.RectPoint(step.K, r.V)
	start := rsFrom.RectPoint(i, r.V)
	var rCopy = ui.NewRect(r.Ctx, start, r.Width(), r.Height(), r.Index, r.V)
	rCopy.DestPoint = dest
	rCopy.OnFinished = func() {
		if s.aName == step.From {
			rCopy.Color = defaults.DefaultColor.BColor
		} else {
			rCopy.Color = defaults.DefaultColor.AColor
		}
	}
	rTo[step.K] = rCopy
}

func (s *Screen) splitA(step *mergesort.Step) {
	s.split(step, "a")
}

func (s *Screen) splitB(step *mergesort.Step) {
	s.split(step, "b")
}

func (s *Screen) split(step *mergesort.Step, name string) {
	var rs *ui.RectSlice
	var color string
	if name == "a" {
		rs = s.rsA
	} else {
		rs = s.rsB
	}
	color = defaults.DefaultColor.CColor

	r := rs.Rectangles
	begin := r[step.IBegin]
	end := r[step.IEnd-1]
	begin.OnDrawing = func() {
		begin.Color = color
	}
	end.OnDrawing = func() {
		end.Color = color
	}

	begin.OnFinished = func() {
		begin.Color = color
	}
	end.OnFinished = func() {
		end.Color = color
	}
}

func (s *Screen) draw(progress float64) {
	s.ctx.FillStyle = defaults.DefaultColor.BackgroundColor
	s.ctx.FillRect(0, 0, float64(s.c.Width), float64(s.c.Height))

	s.rsA.Draw(progress)
	s.rsB.Draw(progress)

	s.ready = s.rsA.Ready() && s.rsB.Ready()
}

func createCanvas(id string, size int) *js.Object {
	body := js.Global.Get("document").Call("getElementById", id)
	obj := js.Global.Get("document").Call("createElement", "canvas")
	obj.Set("width", strconv.Itoa(canvasWidth(size)))
	obj.Set("height", strconv.Itoa(canvasHeight(size, 2)))
	body.Set("innerHTML", "")
	body.Call("appendChild", obj)
	return obj
}

func canvasWidth(size int) int {
	return defaults.BarWidth*size + (size-1)*defaults.BarSpace
}

func canvasHeight(size int, index int) int {
	return index * rectanglesHeight(size, index)
}

func rectanglesHeight(size int, index int) int {
	return size*defaults.HeightUnit + (index-1 )*3*defaults.HeightUnit
}
