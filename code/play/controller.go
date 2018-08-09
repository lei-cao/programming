package play

import (
	"github.com/lei-cao/learning-cs-again/code/visualizer"
	"github.com/lei-cao/learning-cs-again/code/utils"
	"github.com/lei-cao/learning-cs-again/code/algorithms/sorting"
	"github.com/lei-cao/learning-cs-again/code/visualizer/sorting/basic"
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
	animation visualizer.Animator
	config    *ControllerConfig
	sorter    sorting.Sorter
	nums      []int
}

// Update config. Being called by JS
func (c *Controller) UpdateConfig(config *ControllerConfig) {
	c.config = config
	c.animation.UpdateDuration(config.Duration)
}

// Init the visualizer controller
func (c *Controller) Init(config *ControllerConfig) {
	c.config = config

	if config.Size == 0 {
		c.config.Size = defaultSize
	}

	// Init animation
	c.animation = visualizer.NewAnimation()
	c.animation.UpdateDuration(config.Duration)

	// Initial states for screen
	c.nums = utils.Shuffle(c.config.Size)
	s := basic.NewScreen(c.config.Id, c.config.Size, c.nums)
	c.animation.SetScreen(s)

	// Apply algorithms, set steps
	c.applyAlgorithm(config)

	c.animation.StartAnimating()
}

// Stop auto running. Switch to manual control
func (c *Controller) Stop() {
	c.animation.Stop()
}

// Resume auto running
func (c *Controller) Resume() {
	c.animation.Resume()
}

// Do next step from the steps queue
// Being triggered manually or automatically
func (c *Controller) NextStep() {
	c.animation.NextStep()
}

func (c *Controller) applyAlgorithm(config *ControllerConfig) {
	switch config.Id {
	case "bubble":
		c.sorter = sorting.NewBubbleSort()
	case "selection":
		c.sorter = sorting.NewSelectionSort()
	case "insertion":
		c.sorter = sorting.NewInsertionSort()
	case "quick":
		c.sorter = sorting.NewQuickSort()
	case "topDownMergeSort":
		c.sorter = sorting.NewTopDownMergeSort()
	}
	c.sorter.Sort(c.nums)
	c.animation.SetSteps(c.sorter.Steps())
}
