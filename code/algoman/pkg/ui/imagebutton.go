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
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

func NewImageButton(rect image.Rectangle, imgOn *ebiten.Image, imgOff *ebiten.Image) *ImageButton {

	stopBtn := &ImageButton{
		Rect:      rect,
		Img:       imgOff,
		ImgActive: imgOn,
	}
	return stopBtn

}

type ImageButton struct {
	ImgActive *ebiten.Image
	Img       *ebiten.Image
	Rect      image.Rectangle

	mouseDown bool

	onPressed func(b *ImageButton)
}

func (b *ImageButton) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
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

func (b *ImageButton) Draw(dst *ebiten.Image) {
	t := imageTypeButton
	if b.mouseDown {
		t = imageTypeButtonPressed
	}
	op := &ebiten.DrawImageOptions{}
	x, y := b.ImgActive.Size()
	op.GeoM.Scale(float64(defaults.ImageBtnWidth)/float64(x), float64(defaults.ImageBtnHeight)/float64(y))
	op.GeoM.Scale(0.5, 0.5)
	//op.GeoM.Scale(ebiten.DeviceScaleFactor(), ebiten.DeviceScaleFactor())
	op.GeoM.Translate(float64(b.Rect.Min.X), float64(b.Rect.Min.Y))
	if t == imageTypeButtonPressed {
		dst.DrawImage(b.ImgActive, op)
	} else {
		dst.DrawImage(b.Img, op)
	}
}

func (b *ImageButton) SetOnPressed(f func(b *ImageButton)) {
	b.onPressed = f
}
