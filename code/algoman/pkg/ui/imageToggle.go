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
	"bytes"
	"github.com/lei-cao/programming/code/algoman/resources/images"
	"log"
	"github.com/lei-cao/programming/code/algoman/pkg/consts"
)

var (
	playImg  *ebiten.Image
	pauseImg *ebiten.Image
)

func init() {
	play, _, err := image.Decode(bytes.NewReader(images.PLAY_ON_png))
	if err != nil {
		log.Fatal(err)
	}
	playImg, _ = ebiten.NewImageFromImage(play, ebiten.FilterDefault)

	pause, _, err := image.Decode(bytes.NewReader(images.PAUSE_ON_png))
	if err != nil {
		log.Fatal(err)
	}
	pauseImg, _ = ebiten.NewImageFromImage(pause, ebiten.FilterDefault)
}

func NewImageToggle(rect image.Rectangle) *ImageToggle {

	playToggle := &ImageToggle{
		Rect:   rect,
		ImgOff: playImg,
		ImgOn:  pauseImg,
	}
	return playToggle

}

type ImageToggle struct {
	ImgOn  *ebiten.Image
	ImgOff *ebiten.Image
	Rect   image.Rectangle
	On     bool

	mouseDown bool

	onPressed func(b *ImageToggle)
}

func (b *ImageToggle) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x -= consts.ScreenBorder
		y -= consts.ScreenBorder
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
}

func (b *ImageToggle) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.117, 0.117)
	op.GeoM.Translate(float64(b.Rect.Min.X), float64(b.Rect.Min.Y))
	//if b.mouseDown {
	if b.On {
		dst.DrawImage(b.ImgOn, op)
	} else {
		dst.DrawImage(b.ImgOff, op)
	}
	//}
}

func (b *ImageToggle) SetOnPressed(f func(b *ImageToggle)) {
	b.onPressed = f
}
