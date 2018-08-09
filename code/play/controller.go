package play

import (
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
	Animation Animator
	Config    *ControllerConfig
	nums      []int
	numsB     []int
}

// Update config. Being called by JS
func (c *Controller) UpdateConfig(config *ControllerConfig) {
	c.Config = config
	c.Animation.UpdateDuration(config.Duration)
}

// Init the visualizer controller
func (c *Controller) Init(config *ControllerConfig) {
	c.Config = config
	c.Animation = NewAnimation()
	c.Animation.UpdateDuration(config.Duration)

	if config.Size == 0 {
		c.Config.Size = defaultSize
	}

	c.nums = utils.Shuffle(c.Config.Size)
	s := visualizer.NewScreen(c.Config.Id, c.Config.Size, c.nums)
	c.Animation.SetScreen(s)

	c.ApplyAlgorithm(config)

	c.Animation.StartAnimating()
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
	c.Animation.Stop()
}

// Resume auto running
func (c *Controller) Resume() {
	c.Animation.Resume()
}

// Do next step from the steps queue
// Being triggered manually or automatically
func (c *Controller) NextStep() {
	c.Animation.NextStep()
}
