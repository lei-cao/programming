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
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

func NewRectSlice(values []int, offsetX, offsetY float64) *RectSlice {
	rs := new(RectSlice)
	rs.values = values
	rs.offsetX, rs.offsetY = offsetX, offsetY
	for k, v := range rs.values {
		r := NewRectangle(v)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(offsetX+float64((k+1)*(defaults.BarWidth+defaults.BarMargin)), offsetY+float64((len(rs.values)-r.V+2)*defaults.BarHeightUnit))
		//op.SourceRect = r.rect
		r.startOp = op
		r.startIndex = k
		r.endIndex = k

		rs.rectangles = append(rs.rectangles, r)
	}
	return rs
}

type RectSlice struct {
	values     []int
	rectangles []*Rectangle
	aIndex     int
	bIndex     int
	offsetX    float64
	offsetY    float64
}

func (rs *RectSlice) Update(progress float64) {
	for _, r := range rs.rectangles {
		if r.endIndex != r.startIndex {
			r.startOp.GeoM.Reset()
			dx := float64((r.startIndex+1)*(defaults.BarWidth+defaults.BarMargin)) + progress*float64((r.endIndex-r.startIndex)*(defaults.BarWidth+defaults.BarMargin))
			r.startOp.GeoM.Translate(rs.offsetX+dx, rs.offsetY+float64((len(rs.values)-r.V+2)*defaults.BarHeightUnit))
			if progress == 1 {
				r.startIndex = r.endIndex
			}
		} else {
			r.startOp.ColorM.Reset()
		}
	}
}

func (rs *RectSlice) NextStep(s visualizer.Stepper) {
	if step, ok := s.(*basicsort.Step); ok {
		if step.DoSwap() {
			rs.swap(step.A(), step.B())
		} else {
			rs.pass(step.A(), step.B())
		}
	}
}

func (rs *RectSlice) Draw(image *ebiten.Image) {
	for _, r := range rs.rectangles {
		r.isB = false
		r.isA = false
		r.barImage.Fill(defaults.BarColor)
		if rs.aIndex == r.startIndex {
			r.barImage.Fill(defaults.ColorA)
		}

		if rs.bIndex == r.startIndex {
			r.barImage.Fill(defaults.ColorB)
		}
		image.DrawImage(r.barImage, r.startOp)
	}
}

func (rs *RectSlice) swap(ia, ib int) {
	if ia < 0 || ib < 0 || ia > len(rs.rectangles)-1 || ib > len(rs.rectangles)-1 {
		return
	}
	rs.aIndex = ia
	rs.bIndex = ib

	ra := rs.rectangles[ia]
	rb := rs.rectangles[ib]
	ra.endIndex = ib
	rb.endIndex = ia
	ra.isA = true
	rb.isB = true
	rs.rectangles[ia] = rb
	rs.rectangles[ib] = ra
}

func (rs *RectSlice) pass(ia, ib int) {
	if ia < 0 || ib < 0 || ia > len(rs.rectangles)-1 || ib > len(rs.rectangles)-1 {
		return
	}

	rs.rectangles[ia].isA = true
	rs.rectangles[ib].isB = true
	rs.aIndex = ia
	rs.bIndex = ib
}

func rectSliceWidth(size int) int {
	return size*defaults.BarWidth + (size-1)*defaults.BarMargin
}
func rectSliceHeight(size int) int {
	return size * defaults.BarHeightUnit
}
