package sorting_test

import (
	"testing"
	"github.com/lei-cao/programming/code/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/algorithms/sorting/heapsort"
	"github.com/lei-cao/programming/code/algorithms/sorting/mergesort"
	"github.com/lei-cao/programming/code/algorithms/sorting"
	"reflect"
)

type testCase struct {
	in  []int
	out []int
}

var cases = []*testCase{
	{
		in:  []int{5, 4, 3, 2, 6, 7, 8, 9, 6, 10, 0},
		out: []int{0, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10},
	},
	{
		in:  []int{1, 1, 2, -1, -2},
		out: []int{-2, -1, 1, 1, 2},
	},
	{
		in:  []int{-1, 1, 2, 3},
		out: []int{-1, 1, 2, 3},
	},
}

func TestSorters(t *testing.T) {
	sorters := []sorting.Sorter{
		basicsort.NewBubbleSort(),
		basicsort.NewSelectionSort(),
		basicsort.NewInsertionSort(),
		basicsort.NewQuickSort(),
		mergesort.NewTopDownMergeSort(),
		&heapsort.HeapSort{},
	}
	for _, s := range sorters {
		for _, c := range cases {
			var a []int = make([]int, len(c.in))
			copy(a, c.in)

			s.Sort(a)
			if len(a) != len(c.out) {
				t.Fatalf("Unsorted! by %s. in: %v, out: %v", reflect.TypeOf(s), a, c.out)
			}
			for k, v := range c.out {
				if v != a[k] {
					t.Fatalf("Unsorted! by %s. in: %v, out: %v", reflect.TypeOf(s), a, c.out)
				}
			}
		}
	}
}
