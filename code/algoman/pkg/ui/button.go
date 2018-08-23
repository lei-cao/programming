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
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"github.com/lei-cao/programming/code/algoman/resources/images"
	"log"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
	"golang.org/x/image/font/gofont/gomono"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

var (
	uiImage       *ebiten.Image
	uiFont        font.Face
	uiFontMHeight int
)

func init() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.UI_png))
	if err != nil {
		log.Fatal(err)
	}
	uiImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	tt, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}
	uiFont = truetype.NewFace(tt, &truetype.Options{
		Size:    float64(12 * defaults.DeviceScale),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	b, _, _ := uiFont.GlyphBounds('M')
	uiFontMHeight = (b.Max.Y - b.Min.Y).Ceil()
}

type imageType int

const (
	imageTypeButton        imageType = iota
	imageTypeButtonPressed
)

var imageSrcRects = map[imageType]image.Rectangle{
	imageTypeButton:        image.Rect(0, 0, 16, 16),
	imageTypeButtonPressed: image.Rect(16, 0, 32, 16),
}

type Input struct {
	mouseButtonState int
}

func drawNinePatches(dst *ebiten.Image, dstRect image.Rectangle, srcRect image.Rectangle) {
	srcX := srcRect.Min.X
	srcY := srcRect.Min.Y
	srcW := srcRect.Dx()
	srcH := srcRect.Dy()

	dstX := dstRect.Min.X
	dstY := dstRect.Min.Y
	dstW := dstRect.Dx()
	dstH := dstRect.Dy()

	op := &ebiten.DrawImageOptions{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			op.GeoM.Reset()

			sx := srcX
			sy := srcY
			sw := srcW / 4
			sh := srcH / 4
			dx := 0
			dy := 0
			dw := sw
			dh := sh
			switch i {
			case 1:
				sx = srcX + srcW/4
				sw = srcW / 2
				dx = srcW / 4
				dw = dstW - 2*srcW/4
			case 2:
				sx = srcX + 3*srcW/4
				dx = dstW - srcW/4
			}
			switch j {
			case 1:
				sy = srcY + srcH/4
				sh = srcH / 2
				dy = srcH / 4
				dh = dstH - 2*srcH/4
			case 2:
				sy = srcY + 3*srcH/4
				dy = dstH - srcH/4
			}

			op.GeoM.Scale(float64(dw)/float64(sw), float64(dh)/float64(sh))
			op.GeoM.Translate(float64(dx), float64(dy))
			op.GeoM.Translate(float64(dstX), float64(dstY))
			r := image.Rect(sx, sy, sx+sw, sy+sh)
			op.SourceRect = &r
			dst.DrawImage(uiImage, op)
		}
	}
}

func NewButton(rect image.Rectangle, text string) *Button {
	bounds, _ := font.BoundString(uiFont, text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	if w < defaults.ButtonMinWidth {
		w = defaults.ButtonMinWidth
	}
	rect.Max.X = rect.Min.X + w + defaults.ButtonPadding
	btn := &Button{
		Rect: rect,
		Text: text,
	}
	return btn
}

type Button struct {
	Rect image.Rectangle
	Text string

	mouseDown bool

	onPressed func(b *Button)
}

func (b *Button) Update() {
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
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *Button) Draw(dst *ebiten.Image) {
	t := imageTypeButton
	if b.mouseDown {
		t = imageTypeButtonPressed
	}
	drawNinePatches(dst, b.Rect, imageSrcRects[t])

	bounds, _ := font.BoundString(uiFont, b.Text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := b.Rect.Min.X + (b.Rect.Dx()-w)/2
	y := b.Rect.Max.Y - (b.Rect.Dy()-uiFontMHeight)/2
	text.Draw(dst, b.Text, uiFont, x, y, color.Black)
}

func (b *Button) SetOnPressed(f func(b *Button)) {
	b.onPressed = f
}
