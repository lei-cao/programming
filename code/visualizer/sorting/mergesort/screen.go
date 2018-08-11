package mergesort

import (
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/learning-cs-again/code/visualizer"
	"github.com/lei-cao/learning-cs-again/code/algorithms/sorting/mergesort"
	"github.com/lei-cao/learning-cs-again/code/visualizer/defaults"
)

func NewScreen(id string, size int, nums []int) *Screen {
	s := new(Screen)
	s.id = id
	s.size = size

	obj := createCanvas(s.id, s.size)
	s.c = canvas.New(obj)
	s.ctx = s.c.GetContext2D()

	pa := &Point{X: 0, Y: 0}
	s.rsA = NewRectSlice(s.ctx, nums, pa, "a", false)

	pb := &Point{X: 0, Y: s.rsA.Height() + 30}
	s.rsB = NewRectSlice(s.ctx, nums, pb, "b", false)

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
	rsA             *RectSlice
	rsB             *RectSlice
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
			if s.aName == step.From {
				//s.splitA(step)
			}

			if s.bName == step.From {
				//s.splitB(step)
			}
		}

		if step.IsAssignStep() {
			s.assign(step)
		}
	}
}

func (s *Screen) assign(step *mergesort.Step) {
	var rsFrom *RectSlice
	var rsTo *RectSlice
	if s.aName == step.From {
		rsFrom = s.rsA
		rsTo = s.rsB
	} else {
		rsFrom = s.rsB
		rsTo = s.rsA
	}

	rFrom := rsFrom.Rectangles
	rTo := rsTo.Rectangles

	var r *Rectangle
	var rCopy = &Rectangle{}
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
	rCopy.Color = defaults.DefaultColor.CColor
	//rCopy.destPoint = rTo[step.K].startPoint
	rCopy.destPoint = rsTo.rectPoint(step.K, r.V)
	rCopy.startPoint = rsFrom.rectPoint(i, r.V)
	rCopy.Height = r.Height
	rCopy.Width = r.Width
	rCopy.Ctx = r.Ctx
	rCopy.V = r.V
	r.Color = defaults.DefaultColor.AColor
	rTo[step.K] = rCopy
}

func (s *Screen) splitA(step *mergesort.Step) {
	s.split(step, "a")
}

func (s *Screen) splitB(step *mergesort.Step) {
	s.split(step, "b")
}

func (s *Screen) split(step *mergesort.Step, name string) {
	var rs *RectSlice
	if name == "a" {
		rs = s.rsA
	} else {
		rs = s.rsB
	}

	r := rs.Rectangles
	begin := r[step.IBegin]
	mid := r[step.IMid]
	end := r[step.IEnd]
	begin.Color = defaults.DefaultColor.AColor
	mid.Color = defaults.DefaultColor.CColor
	end.Color = defaults.DefaultColor.AColor
	r[step.IBegin] = begin
	r[step.IMid] = mid
	r[step.IEnd] = end
}

func (s *Screen) draw(progress float64) {
	s.ctx.FillStyle = defaults.DefaultColor.BackgroundColor
	s.ctx.FillRect(0, 0, float64(s.c.Width), float64(s.c.Height))
	//s.ctx.FillStyle = defaults.DefaultColor.BarColor
	//s.ctx.Font = "20px Arial";
	//s.ctx.FillText(s.aName, 5, 20, -1);
	for k, r := range s.rsA.Rectangles {
		s.finishedDrawing[k] = r.Animate(progress)
	}
	for k, r := range s.rsB.Rectangles {
		s.finishedDrawing[s.size+k] = r.Animate(progress)
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
	obj.Set("height", strconv.Itoa(canvasHeight(size, 2)+60))
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
