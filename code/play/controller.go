package controller

import (
	"github.com/lei-cao/learning-cs-again/code/sort"
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-canvas"
	"time"
)

var barWidth = 8
var barSpace = 2
var heightUnit = 5
// default size for the slice being solved
var defaultSize = 10
var velocity float64

// Hold the operation steps queue
type Step struct {
	A      int
	B      int
	DoSwap bool
	Next   *Step
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
	startTime      int
	now            int
	then           int
	elapsed        int
}

type ControllerConfig struct {
	Velocity float64 `json:"velocity"`
	Size     int     `json:"size"`
}

func (c *ControllerConfig) SetVelocity(v float64) {
	c.Velocity = v
	velocity = c.Velocity
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
	c.fps = 60
	// 10 pixel per second
	velocity = config.Velocity
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
			step.B = j + 1
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				step.DoSwap = true
			}
			c.LastStep.Next = step
			c.LastStep = step
			numSteps ++
		}
	}

	obj := createCanvas(id, c.Config.Size)
	c.C = canvas.New(obj)
	c.Ctx = c.C.GetContext2D()
	c.startAnimating()
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
	c.CurrentStep = c.CurrentStep.Next
	c.update()
}

// Update the screen states based on current step
func (c *Controller) update() {
	a := c.Screen.Rectangles[c.CurrentStep.A]
	b := c.Screen.Rectangles[c.CurrentStep.B]
	a.IsA = true
	b.IsB = true
	a.Waiting = 2
	//b.Waiting = 1
	c.Screen.AIndex = c.CurrentStep.A
	c.Screen.BIndex = c.CurrentStep.B
	if c.CurrentStep.DoSwap {
		a.ToIndex = c.CurrentStep.B
		b.ToIndex = c.CurrentStep.A

		c.Screen.Rectangles[c.CurrentStep.A] = b
		c.Screen.Rectangles[c.CurrentStep.B] = a
	}
}

// Draw the screen
func (c *Controller) draw() {
	c.Screen.draw(c.Ctx, c.fpdInterval)
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
	c.then = int(time.Now().UnixNano())
	c.startTime = c.then
	c.animate()
}

func (c *Controller) animate() {
	c.animationFrame = js.Global.Call("requestAnimationFrame", c.animate)
	c.now = int(time.Now().UnixNano())
	c.elapsed = c.now - c.then
	if c.elapsed > int(c.fpdInterval) {
		c.then = c.now - (c.elapsed % int(c.fpdInterval))

		c.Animating = true
		c.Ctx.ClearRect(0, 0, float64(canvasWidth(c.Config.Size)), float64(canvasHeight(c.Config.Size)))
		c.draw()
	}
}
