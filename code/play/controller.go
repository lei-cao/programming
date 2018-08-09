package play

import (
	"github.com/gopherjs/gopherjs/js"
	"time"
	"math"
	"github.com/lei-cao/learning-cs-again/code/visualizer"
	"github.com/lei-cao/learning-cs-again/code/utils"
)

// default size for the slice being solved
var defaultSize = 10


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
	Steps          visualizer.Stepper
	Screen         visualizer.Screener
	AutoUpdate     bool
	Animating      bool
	Config         *ControllerConfig
	animationFrame *js.Object
	fps            int
	fpdInterval    float64
	startTime      float64
	now            float64
	then           float64
	elapsed        float64
	nums           []int
	numsB          []int
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
	c.Steps = visualizer.NewStep()
	c.Duration = config.Duration
	c.fps = 60

	if config.Size == 0 {
		c.Config.Size = defaultSize
	}

	c.nums = utils.Shuffle(c.Config.Size)
	s := visualizer.NewScreen(c.Config.Id, c.Config.Size, c.nums)
	c.Screen = s

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
	case "topDownMergeSort":
		c.TopDownMergeSort()
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
	if c.Steps.Finished() {
		js.Global.Call("cancelAnimationFrame", c.animationFrame)
		return
	}

	if !c.Animating {
		c.startAnimating()
	}
	if !c.Screen.Ready() {
		return
	}
	c.startTime = makeTimestamp()
	c.update()
}

// Update the screen states based on current step
func (c *Controller) update() {
	c.Screen.Update(c.Steps.NextStep())
}

// Draw the screen
func (c *Controller) draw(timestamp float64) {
	c.Screen.Draw(timestamp)
	if c.Screen.Ready() {
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

		c.Screen.Clear()
		c.draw(c.Timing(timeFraction))
	}
}

func makeTimestamp() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Millisecond)
}
