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
	"github.com/fogleman/gg"
	"bytes"
	"log"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"image/color"
	"strconv"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font"
)

var fontSize = 28
var trueFont *truetype.Font
var face font.Face
func NewCircle(value int) *Circle {
	c := new(Circle)

	dc := gg.NewContext(defaults.CircleRadius*2, defaults.CircleRadius*2)
	dc.DrawCircle(float64(defaults.CircleRadius), float64(defaults.CircleRadius), float64(defaults.CircleRadius))
	dc.SetRGB(99, 33, 11)
	dc.Fill()
	dc.SetColor(color.Black)

	var err error
	trueFont, err = truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face = truetype.NewFace(trueFont, &truetype.Options{Size: float64(fontSize)})


	dc.SetFontFace(face)
	dc.DrawString(strconv.Itoa(value), float64(defaults.CircleRadius - fontSize/4), float64(defaults.CircleRadius + fontSize/4))

	var writer = new(bytes.Buffer)

	dc.EncodePNG(writer)

	img, _, err := image.Decode(bytes.NewReader(writer.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	uiImage, err := ebiten.NewImageFromImage(img, ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	c.CircleImage = uiImage

	return c
}

type Circle struct {
	CircleImage *ebiten.Image
	startOp     *ebiten.DrawImageOptions
	startIndex  int
	endIndex    int
	isA         bool
	isB         bool
	rect        *image.Rectangle
	V           int
}
