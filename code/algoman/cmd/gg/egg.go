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

package gg

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten"
	"bytes"
	"image"
	"log"
	"image/color"
)

func NewEbitenGG(img *ebiten.Image) *EbitenGG {
	egg := new(EbitenGG)
	egg.img = img
	egg.initContext()
	return egg
}

type EbitenGG struct {
	dc  *gg.Context
	img *ebiten.Image
}

func (egg *EbitenGG) FlushImage() *ebiten.Image {
	var writer = new(bytes.Buffer)

	egg.dc.EncodePNG(writer)

	ggImg, _, err := image.Decode(bytes.NewReader(writer.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	uiImage, err := ebiten.NewImageFromImage(ggImg, ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	egg.img.DrawImage(uiImage, nil)

	return egg.img
}

func (egg *EbitenGG) Image() *ebiten.Image {
	return egg.img
}

func (egg *EbitenGG) initContext() {
	x, y := egg.img.Size()
	egg.dc = gg.NewContext(x, y)
	egg.dc.SetColor(color.White)
	egg.dc.SetLineWidth(5)

}
