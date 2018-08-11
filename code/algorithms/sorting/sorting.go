package sorting

import (
	"github.com/lei-cao/learning-cs-again/code/visualizer"
)

type Sorter interface {
	Sort(nums []int)
	Steps() visualizer.Stepper
}
