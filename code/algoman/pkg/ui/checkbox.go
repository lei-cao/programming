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

package ui

import (
	"image"
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

const (
	imageTypeCheckBox        imageType = iota
	imageTypeCheckBoxPressed
	imageTypeCheckBoxMark
)

var cbImageSrcRects = map[imageType]image.Rectangle{
	imageTypeCheckBox:        image.Rect(0, 32, 16, 48),
	imageTypeCheckBoxPressed: image.Rect(16, 32, 32, 48),
	imageTypeCheckBoxMark:    image.Rect(32, 32, 48, 48),
}

const (
	checkboxWidth       = 16
	checkboxHeight      = 16
	checkboxPaddingLeft = 8
)

type CheckBox struct {
	X    int
	Y    int
	Text string

	checked   bool
	value     string
	mouseDown bool

	onCheckChanged func(c *CheckBox)
}

func (c *CheckBox) width() int {
	b, _ := font.BoundString(uiFont, c.Text)
	w := (b.Max.X - b.Min.X).Ceil()
	return checkboxWidth + checkboxPaddingLeft + w
}

func (c *CheckBox) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if c.X <= x && x < c.X+c.width() && c.Y <= y && y < c.Y+checkboxHeight {
			c.mouseDown = true
		} else {
			c.mouseDown = false
		}
	} else {
		if c.mouseDown {
			c.checked = !c.checked
			if c.onCheckChanged != nil {
				c.onCheckChanged(c)
			}
		}
		c.mouseDown = false
	}

	if len(ebiten.TouchIDs()) > 0 {
		for _, t := range ebiten.TouchIDs() {
			x, y := ebiten.TouchPosition(t)
			if c.X <= x && x < c.X+c.width() && c.Y <= y && y < c.Y+checkboxHeight {
				c.mouseDown = true
			} else {
				c.mouseDown = false
			}

		}
	} else {
		if c.mouseDown {
			c.checked = !c.checked
			if c.onCheckChanged != nil {
				c.onCheckChanged(c)
			}
		}
		c.mouseDown = false
	}
}

func (c *CheckBox) Draw(dst *ebiten.Image) {
	t := imageTypeCheckBox
	if c.mouseDown {
		t = imageTypeCheckBoxPressed
	}
	r := image.Rect(c.X, c.Y, c.X+checkboxWidth, c.Y+checkboxHeight)
	drawNinePatches(dst, r, cbImageSrcRects[t])
	if c.checked {
		drawNinePatches(dst, r, cbImageSrcRects[imageTypeCheckBoxMark])
	}

	x := c.X + checkboxWidth + checkboxPaddingLeft
	y := (c.Y + 16) - (16-uiFontMHeight)/2
	text.Draw(dst, c.Text, uiFont, x, y, defaults.BarColor)
}

func (c *CheckBox) Checked() bool {
	return c.checked
}

func (c *CheckBox) Check() {
	c.checked = true
}

func (c *CheckBox) UnCheck() {
	c.checked = false
}

func (c *CheckBox) Value() string {
	return c.value
}

func (c *CheckBox) SetValue(v string) {
	c.value = v
}

func (c *CheckBox) SetOnCheckChanged(f func(c *CheckBox)) {
	c.onCheckChanged = f
}
