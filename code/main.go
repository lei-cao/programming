package main

import (
	"github.com/lei-cao/learn/sort"
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

var nums = sort.Shuffle(30)

func main() {
	//js.Module.Get("exports").Set("Algorithm", Algorithm)

	js.Global.Set("algorithm", map[string]interface{}{
		"Algorithm": Algorithm,
	})
}

type Visualizer struct {

}

func (v *Visualizer) Show() {
	go func() {

		sorters := [] sort.Sort{
			new(sort.BubbleSort),
			new(sort.BubbleSort),
			new(sort.BubbleSortSwapped),
		}
		ch := make(chan bool, len(sorters))
		for _, sorter := range sorters {
			go sort.DoSort(nums, sorter, ch)
		}

		searchers := [] sort.Search{
			new(sort.DefaultSearch),
		}
		sch := make(chan int, len(searchers))
		for _, searcher := range searchers {
			go sort.DoSearch(searcher, 14, nums, sch)
		}
		for _ = range sorters {
			fmt.Println(<-ch)
		}
		for _ = range searchers {
			fmt.Println(<-sch)
		}
		close(ch)
		close(sch)

	}()
}


func Algorithm() *js.Object{
	return js.MakeWrapper(&Visualizer{})
}


