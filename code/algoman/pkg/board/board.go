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
)

func init() {
}

func NewBoard(values []int) *Board {
	b := &Board{values: values}
	b.rs = shapes.NewRectSlice(b.values)
	b.BoardImage, _ = ebiten.NewImage(defaults.ScreenWidth, defaults.ScreenHeight, ebiten.FilterDefault)
	b.BoardImage.Fill(defaults.BackgroundColor)
	return b
}

type Board struct {
	BoardImage *ebiten.Image
	Finished   bool
	progress   float64
	values     []int
	rs         *shapes.RectSlice
}

func (b *Board) Draw() {
	b.BoardImage.Clear()
	b.BoardImage.Fill(defaults.BackgroundColor)
	b.rs.Draw(b.BoardImage)
}

func (b *Board) Update(progress float64) {
	b.progress = progress
	b.rs.Update(progress)
}

func (b *Board) NextStep(step visualizer.Stepper) {
	b.rs.NextStep(step)
}

func (b *Board) Ready() bool {
	if b.progress == 1 {
		return true
	}
	return false
}
