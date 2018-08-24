// Copyright 2018 The Algoman Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package board

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/pkg/shapes"
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
)

func init() {
}

func NewSortingBoard(values []int) *SortingBoard {
	b := &SortingBoard{values: values}
	b.rs = shapes.NewRectSlice(b.values)
	b.image, _ = ebiten.NewImage(defaults.ScreenWidth, defaults.ScreenHeight, ebiten.FilterDefault)
	b.image.Fill(defaults.BackgroundColor)
	return b
}

type SortingBoard struct {
	image    *ebiten.Image
	Finished bool
	progress float64
	values   []int
	rs       *shapes.RectSlice
	sorter   sorting.Sorter
}

func (b *SortingBoard) Draw() {
	b.image.Clear()
	b.image.Fill(defaults.BackgroundColor)
	b.rs.Draw(b.image)
}

func (b *SortingBoard) Update(progress float64) {
	b.progress = progress
	b.rs.Update(progress)
}

func (b *SortingBoard) NextStep(step visualizer.Stepper) {
	b.rs.NextStep(step)
}

func (b *SortingBoard) Ready() bool {
	if b.progress == 1 {
		return true
	}
	return false
}

func (b *SortingBoard) Image() *ebiten.Image {
	return b.image
}

func (b *SortingBoard) Steps(id string) visualizer.Stepper {
	switch id {
	case "bubble":
		b.sorter = basicsort.NewBubbleSort()
	case "selection":
		b.sorter = basicsort.NewSelectionSort()
	case "insertion":
		b.sorter = basicsort.NewInsertionSort()
	case "quick":
		b.sorter = basicsort.NewQuickSort()
	case "heap":
		b.sorter = basicsort.NewHeapSort()
		//case "topDownMergeSort":
		//	s := merge.NewScreen(c.config.Id, c.config.Size, c.nums)
		//	c.animation.SetScreen(s)
		//	c.sorter = mergesort.NewTopDownMergeSort()
	}
	b.sorter.Sort(b.values)
	return b.sorter.Steps()
}
