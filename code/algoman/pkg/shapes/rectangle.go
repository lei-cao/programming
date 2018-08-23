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

package shapes

import (
	"image"
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

func NewRectangle(value int) *Rectangle {
	r := new(Rectangle)
	rect := image.Rect(0, 0, barWidth, value*barHeightUnit)
	r.barImage, _ = ebiten.NewImage(barWidth, value*barHeightUnit, ebiten.FilterDefault)
	r.barImage.Fill(defaults.BarColor)
	r.rect = &rect
	r.V = value
	return r
}

type Rectangle struct {
	barImage   *ebiten.Image
	startOp    *ebiten.DrawImageOptions
	startIndex int
	endIndex   int
	isA        bool
	isB        bool
	rect       *image.Rectangle
	V          int
}
