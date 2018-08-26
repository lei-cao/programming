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
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting"
)

func init() {
}

func NewHeapBoard(values []int) *HeapBoard {
	b := &HeapBoard{values: values}
	b.tree, b.dts = shapes.NewDrawTree(b.values)
	image, _ := ebiten.NewImage(defaults.ScreenWidth, defaults.ScreenHeight, ebiten.FilterDefault)
	image.Fill(defaults.BackgroundColor)
	b.egg = shapes.NewEbitenGG(image)
	b.tree.Draw(b.egg, 0)
	b.dts.Draw(b.egg)
	b.egg.FlushImage()
	return b
}

type HeapBoard struct {
	egg      *shapes.EbitenGG
	Finished bool
	progress float64
	values   []int
	tree     *shapes.DrawTree
	dts      *shapes.DrawTreeSlice
	sorter   sorting.Sorter
}

func (b *HeapBoard) Draw() {
	b.tree.Draw(b.egg, 0)
	b.dts.Draw(b.egg)
	b.egg.FlushImage()
}

func (b *HeapBoard) Update(progress float64) {
	b.progress = progress
	b.dts.Update(progress)
}

func (b *HeapBoard) NextStep(step visualizer.Stepper) {
	b.dts.NextStep(step)
}

func (b *HeapBoard) Ready() bool {
	if b.progress == 1 {
		return true
	}
	return false
}

func (b *HeapBoard) Image() *ebiten.Image {
	return b.egg.Image()
}

func (b *HeapBoard) Steps(id string) visualizer.Stepper {
	b.sorter = basicsort.NewHeapSort()
	b.sorter.Sort(b.values)
	return b.sorter.Steps()
}
