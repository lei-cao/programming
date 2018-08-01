package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lei-cao/learning-cs-again/code/sort"
)

var nums = sort.Shuffle(30)

var visualizers map[string]sort.Sort

func main() {
	visualizers = map[string]sort.Sort{}
	visualizers["bubble"] = new(sort.BubbleSort)
	visualizers["bubble_swapped"] = new(sort.BubbleSortSwapped)
	visualizers["selection"] = new(sort.SelectionSort)

	js.Global.Set("algorithm", map[string]interface{}{
		"Algorithm": Algorithm,
	})
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

func Algorithm() *js.Object {
	return js.MakeWrapper(new(Visualizer))
}
