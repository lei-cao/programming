package sorting

import (
	"github.com/lei-cao/programming/code/visualizer"
)

type Sorter interface {
	Sort(a []int)
	Steps() visualizer.Stepper
}
