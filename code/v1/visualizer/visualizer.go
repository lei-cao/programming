package visualizer

import (
	"github.com/lei-cao/programming/code/v1/sort"
	"github.com/lei-cao/programming/code/utils"
)

var nums = utils.Shuffle(30)

var visualizers = map[string]sort.Sort{
	"bubble":         new(sort.BubbleSort),
	"bubble_swapped": new(sort.BubbleSortSwapped),
	"selection":      new(sort.SelectionSort),
}

type Visualizer struct {
}

func (v *Visualizer) Display(id string) {
	go func() {
		if s, ok := visualizers[id]; ok {
			sort.DoSort(nums, s, id)
		}
	}()
}
