package controller

import (
	"github.com/lei-cao/learning-cs-again/code/sort"
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
)

var barWidth = 8
var barSpace = 2
var heightUnit = 5
// default size for the slice being solved
var defaultSize = 10

// Hold the operation steps queue
type Step struct {
	A    int
	B    int
	Next *Step
}

// The visualizer controller
type Controller struct {
	C              *canvas.Canvas
	Ctx            *canvas.Context2D
	Steps          *Step
	LastStep       *Step
	CurrentStep    *Step
	AutoUpdate     bool
	Animating      bool
	Config         *ControllerConfig
	Screen         *Screen
	animationFrame *js.Object
	StopChan       chan bool
}

type ControllerConfig struct {
	Speed int `json:"speed"`
	Size  int `json:"size"`
}

func (c *ControllerConfig) SetSpeed(speed int) {
	c.Speed = speed
}

func (c *ControllerConfig) SetSize(size int) {
	c.Size = size
}

// Update config. Being called by JS
func (c *Controller) UpdateConfig(config *ControllerConfig) {
	c.Config = config
}

// Init the visualizer controller
func (c *Controller) Init(id string, config *ControllerConfig) {
	c.Config = config
	c.AutoUpdate = true
	c.Steps = &Step{}
	c.LastStep = c.Steps
	c.CurrentStep = c.Steps
	c.Screen = new(Screen)
	c.StopChan = make(chan bool)
	if config.Size == 0 {
		c.Config.Size = defaultSize
	}

	c.Screen.Rectangles = []*Rectangle{}
	var nums = sort.Shuffle(c.Config.Size)

	for k, v := range nums {
		r := NewRect(c.Config.Size, k, v)
		c.Screen.Rectangles = append(c.Screen.Rectangles, r)
	}
	c.Screen.FinishedDrawing = map[int]bool{}
	var numSteps int
	for i := 0; i < c.Config.Size; i++ {
		for j := 0; j < c.Config.Size-1; j++ {
			step := &Step{}
			step.A = j
			step.B = j
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				step.B = j + 1
			}
			c.LastStep.Next = step
			c.LastStep = step
			numSteps ++
		}
	}

	obj := createCanvas(id, c.Config.Size)
	c.C = canvas.New(obj)
	c.Ctx = c.C.GetContext2D()
	c.animate()
}

// Stop auto running. Switch to manual control
func (c *Controller) Stop() {
	js.Global.Call("cancelAnimationFrame", c.animationFrame)
	c.Animating = false
	c.AutoUpdate = false
}

// Resume auto running
func (c *Controller) Resume() {
	if !c.AutoUpdate {
		c.AutoUpdate = true
		c.animate()
	}
}

// Do next step from the steps queue
// Being triggered manually or automatically
func (c *Controller) NextStep() {
	if c.CurrentStep.Next == nil {
		js.Global.Call("cancelAnimationFrame", c.animationFrame)
		return
	}
	if !c.Animating {
		c.animate()
	}
	c.CurrentStep = c.CurrentStep.Next
	c.update()
}

// Update the screen states based on current step
func (c *Controller) update() {
	a := c.Screen.Rectangles[c.CurrentStep.A]
	b := c.Screen.Rectangles[c.CurrentStep.B]
	a.IsA = true
	b.IsB = true
	a.ToIndex = c.CurrentStep.B
	b.ToIndex = c.CurrentStep.A

	c.Screen.Rectangles[c.CurrentStep.A] = b
	c.Screen.Rectangles[c.CurrentStep.B] = a
}

// Draw the screen
func (c *Controller) draw() {
	c.Screen.draw(c.Ctx)
	if c.Screen.Ready {
		if c.AutoUpdate {
			c.NextStep()
		} else {
			// Not auto update, cancelAnimationFrame
			c.Stop()
		}
	}
}

func (c *Controller) animate() {
	c.animationFrame = js.Global.Call("requestAnimationFrame", c.animate)
	c.Animating = true
	c.Ctx.ClearRect(0, 0, float64(canvasWidth(c.Config.Size)), float64(canvasHeight(c.Config.Size)))
	c.draw()
}

// The screen including the elements on the canvas
// Maintain the ready for next state
type Screen struct {
	Rectangles      []*Rectangle
	FinishedDrawing map[int]bool
	Ready           bool
}

func (s *Screen) draw(ctx *canvas.Context2D) {
	for k, r := range s.Rectangles {
		s.FinishedDrawing[k] = r.Draw(ctx)
	}
	for _, finished := range s.FinishedDrawing {
		if !finished {
			s.Ready = false
			return
		}
	}
	s.Ready = true
}

// Represent the element in the problem slice
type Rectangle struct {
	Index   int
	ToIndex int
	IsA     bool
	IsB     bool
	Total   int
	Value   int
	Left    float64
	Top     float64
	Width   float64
	Height  float64
	Moving  bool
}

func (r *Rectangle) Draw(ctx *canvas.Context2D) bool {
	var finished bool
	if r.Left != r.toLeft() {
		r.Moving = true
		if r.Index > r.ToIndex {
			r.Left -= 2
		} else if r.Index < r.ToIndex {
			r.Left += 2
		}
	} else {
		r.Index = r.ToIndex
		r.IsA = false
		r.IsB = false
		finished = true
	}

	r.draw(ctx)
	return finished
}

func (r *Rectangle) draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "#B9314F"
	if r.IsA {
		ctx.FillStyle = "#2E86AB"
	}
	if r.IsB {
		ctx.FillStyle = "#12355B"
	}
	ctx.FillRect(r.Left, r.Top, r.Width, r.Height)
}

func (r *Rectangle) toLeft() float64 {
	return float64((barWidth + barSpace) * r.ToIndex)
}

func NewRect(total int, index int, value int) *Rectangle {
	r := new(Rectangle)
	r.Index = index
	r.ToIndex = index
	r.Total = total
	r.Value = value
	r.Left = float64((barWidth + barSpace) * index)
	r.Top = float64(canvasHeight(total) - value*heightUnit)
	r.Width = float64(barWidth)
	r.Height = float64(value * heightUnit)
	return r
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
