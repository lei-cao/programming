package sorting

import (
	"github.com/lei-cao/programming/code/v2/visualizer"
)

type Sorter interface {
	Sort(a []int)
	Steps() visualizer.Stepper
}
