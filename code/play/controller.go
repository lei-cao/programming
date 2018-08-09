package controller

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-canvas"
	"time"
	"math"
	"github.com/lei-cao/learning-cs-again/code/sort"
)

// default size for the slice being solved
var defaultSize = 10

// Hold the operation steps queue
type Step struct {
	A      int
	B      int
	DoSwap bool
	Next   *Step
}

type ColorScheme struct {
	BackgroundColor string
	BarColor        string
	AColor          string
	BColor          string
}

var defaultColor = ColorScheme{
	"#012A36",
	"#A8A7A0",
	"#2AB7B7",
	"#0E7C7B",
}

type ControllerConfig struct {
	Id       string  `json:"id"`
	Duration float64 `json:"duration"`
	Size     int     `json:"size"`
}

func (c *ControllerConfig) SetDuration(s float64) {
	c.Duration = s
}

func (c *ControllerConfig) SetSize(size int) {
	c.Size = size
}

func (c *ControllerConfig) SetId(id string) {
	c.Id = id
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
	fps            int
	fpdInterval    float64
	startTime      float64
	now            float64
	then           float64
	elapsed        float64
	nums           []int
	numsB           []int
	Duration       float64
	Timing         func(timeFraction float64) float64
}

// Update config. Being called by JS
func (c *Controller) UpdateConfig(config *ControllerConfig) {
	c.Config = config
	c.Duration = config.Duration
}

// Init the visualizer controller
func (c *Controller) Init(config *ControllerConfig) {
	c.Config = config
	c.AutoUpdate = true
	c.Steps = &Step{}
	c.LastStep = c.Steps
	c.CurrentStep = c.Steps
	c.Screen = new(Screen)
	c.Duration = config.Duration
	c.fps = 60
	obj := createCanvas(c.Config.Id, c.Config.Size)
	c.C = canvas.New(obj)
	c.Ctx = c.C.GetContext2D()
	c.Screen.C = c.C
	c.Screen.Ctx = c.Ctx
	if config.Size == 0 {
		c.Config.Size = defaultSize
	}

	c.Screen.Rectangles = []*Rectangle{}
	c.nums = sort.Shuffle(c.Config.Size)

	for k, v := range c.nums {
		r := NewRect(c.Config.Size, k, v, c.Ctx)
		c.Screen.Rectangles = append(c.Screen.Rectangles, r)
	}
	c.Screen.FinishedDrawing = map[int]bool{}

	c.ApplyAlgorithm(config)

	c.startAnimating()
}

func (c *Controller) ApplyAlgorithm(config *ControllerConfig) {
	switch config.Id {
	case "bubble":
		c.BubbleSort()
	case "selection":
		c.SelectionSort()
	case "insertion":
		c.InsertionSort()
	case "quick":
		c.QuickSort()
	}
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
		c.startAnimating()
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
		c.startAnimating()
	}
	if !c.Screen.Ready {
		return
	}
	c.startTime = makeTimestamp()
	c.CurrentStep = c.CurrentStep.Next
	c.update()
}

// Update the screen states based on current step
func (c *Controller) update() {
	if c.CurrentStep.DoSwap {
		c.Screen.swap(c.CurrentStep.A, c.CurrentStep.B)
	} else {
		c.Screen.pass(c.CurrentStep.A, c.CurrentStep.B)
	}
}

// Draw the screen
func (c *Controller) draw(timestamp float64) {
	c.Screen.Draw(timestamp)
	if c.Screen.Ready {
		if c.AutoUpdate {
			c.NextStep()
		} else {
			// Not auto update, cancelAnimationFrame
			c.Stop()
		}
	}
}

func (c *Controller) startAnimating() {
	c.fpdInterval = 1000 / float64(c.fps)
	c.then = makeTimestamp()
	c.startTime = c.then

	c.Timing = func(timeFraction float64) float64 {
		var x = 0.5
		return math.Pow(timeFraction, 2) * ((x+1)*timeFraction - x)
	}
	c.animationFrame = js.Global.Call("requestAnimationFrame", c.animate)
}

func (c *Controller) animate(timestamp float64) {
	c.animationFrame = js.Global.Call("requestAnimationFrame", c.animate)
	c.now = makeTimestamp()
	c.elapsed = c.now - c.then
	c.Animating = true

	if c.elapsed > c.fpdInterval {

		var timeFraction = float64(c.now-c.startTime) / c.Duration
		if timeFraction > 1 {
			timeFraction = 1
		}
		c.then = c.now - math.Mod(c.elapsed, c.fpdInterval)

		c.Ctx.ClearRect(0, 0, float64(canvasWidth(c.Config.Size)), float64(canvasHeight(c.Config.Size)))
		c.draw(c.Timing(timeFraction))
	}
}

func makeTimestamp() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Millisecond)
}
