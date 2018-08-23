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
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

func NewToggleButton(rect image.Rectangle, textOn string, textOff string, isOn bool) *ToggleButton {
	var text = textOn
	if len(textOff) > len(textOn) {
		text = textOff
	}
	bounds, _ := font.BoundString(uiFont, text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	rect.Max.X = rect.Min.X + w + defaults.ButtonPadding
	btn := &ToggleButton{
		Rect:    rect,
		TextOn:  textOn,
		TextOff: textOff,
		On:      isOn,
	}
	return btn
}

type ToggleButton struct {
	Rect    image.Rectangle
	TextOn  string
	TextOff string
	On      bool

	mouseDown bool

	onPressed func(b *ToggleButton)
}

func (b *ToggleButton) Update() {
	if len(ebiten.TouchIDs()) > 0 {
		for _, t := range ebiten.TouchIDs() {
			x, y := ebiten.TouchPosition(t)
			if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
				b.mouseDown = true
			} else {
				b.mouseDown = false
			}
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.On = !b.On
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}

	/*
	if len(ebiten.TouchIDs()) > 0 {
		for _, t := range ebiten.TouchIDs() {
			x, y := ebiten.TouchPosition(t)
			if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
				b.mouseDown = true
			} else {
				b.mouseDown = false
			}
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
	*/
}

func (b *ToggleButton) Draw(dst *ebiten.Image) {
	t := imageTypeButton
	if b.mouseDown {
		t = imageTypeButtonPressed
	}
	drawNinePatches(dst, b.Rect, imageSrcRects[t])
	var currentText string
	if b.On {
		currentText = b.TextOn
	} else {
		currentText = b.TextOff
	}

	bounds, _ := font.BoundString(uiFont, currentText)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := b.Rect.Min.X + (b.Rect.Dx()-w)/2
	y := b.Rect.Max.Y - (b.Rect.Dy()-uiFontMHeight)/2
	text.Draw(dst, currentText, uiFont, x, y, color.Black)
}

func (b *ToggleButton) SetOnPressed(f func(b *ToggleButton)) {
	b.onPressed = f
}
