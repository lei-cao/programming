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

package game

import (
	"image"
	"github.com/lei-cao/programming/code/algoman/pkg/ui"
	"github.com/hajimehoshi/ebiten"
	"bytes"
	"log"
	"github.com/lei-cao/programming/code/algoman/resources/images"
)

var (
	forwardOffImg      *ebiten.Image
	forwardOnImg       *ebiten.Image
	backwardOffImg     *ebiten.Image
	backwardOnImg      *ebiten.Image
	stepForwardOffImg  *ebiten.Image
	stepForwardOnImg   *ebiten.Image
	stepBackwardOffImg *ebiten.Image
	stepBackwardOnImg  *ebiten.Image
	redoOffImg         *ebiten.Image
	redoOnImg          *ebiten.Image
)

var (
	controlY0 = 175
	controlY1 = 190
)

func init() {
	forwardOffImg = imgFromByte(images.FORWARD_OFF_png)
	forwardOnImg = imgFromByte(images.FORWARD_ON_png)
	backwardOffImg = imgFromByte(images.BACKWARD_OFF_png)
	backwardOnImg = imgFromByte(images.BACKWARD_ON_png)
	stepForwardOffImg = imgFromByte(images.STEP_FORWARD_OFF_png)
	stepForwardOnImg = imgFromByte(images.STEP_FORWARD_ON_png)
	stepBackwardOffImg = imgFromByte(images.STEP_BACKWARD_OFF_png)
	stepBackwardOnImg = imgFromByte(images.STEP_BACKWARD_ON_png)
	redoOffImg = imgFromByte(images.REDO_OFF_png)
	redoOnImg = imgFromByte(images.REDO_ON_png)
}

func imgFromByte(imgByte []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImg, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	return ebitenImg
}

func NewController() *Controller {
	c := new(Controller)
	c.PlayToggle = ui.NewImageToggle(image.Rect(10, controlY0, 25, controlY1))
	c.NextStepBtn = ui.NewImageButton(image.Rect(35, controlY0, 50, controlY1), stepForwardOffImg, stepForwardOnImg)
	c.SpeedDownBtn = ui.NewImageButton(image.Rect(60, controlY0, 75, controlY1), backwardOffImg, backwardOnImg)
	c.SpeedUpBtn = ui.NewImageButton(image.Rect(85, controlY0, 100, controlY1), forwardOffImg, forwardOnImg)
	c.RestartBtn = ui.NewImageButton(image.Rect(110, controlY0, 125, controlY1), redoOffImg, redoOnImg)

	c.QuickSortCB = &ui.CheckBox{
		X:    10,
		Y:    controlY1 + 10,
		Text: "Quick Sort",
	}

	c.QuickSortCB.SetValue("quick")

	c.HeapSortCB = &ui.CheckBox{
		X:    10,
		Y:    controlY1 + 10 + 25,
		Text: "Heap Sort",
	}
	c.HeapSortCB.SetValue("heap")

	return c
}

type Controller struct {
	PlayToggle   *ui.ImageToggle
	NextStepBtn  *ui.ImageButton
	SpeedUpBtn   *ui.ImageButton
	SpeedDownBtn *ui.ImageButton
	RestartBtn   *ui.ImageButton

	QuickSortCB *ui.CheckBox
	HeapSortCB *ui.CheckBox
}

func (c *Controller) Update() {
	c.PlayToggle.Update()
	c.NextStepBtn.Update()
	c.SpeedDownBtn.Update()
	c.SpeedUpBtn.Update()
	c.RestartBtn.Update()
	c.QuickSortCB.Update()
	c.HeapSortCB.Update()
}

func (c *Controller) Draw(image *ebiten.Image) {
	c.PlayToggle.Draw(image)
	c.NextStepBtn.Draw(image)
	c.SpeedDownBtn.Draw(image)
	c.SpeedUpBtn.Draw(image)
	c.RestartBtn.Draw(image)
	c.QuickSortCB.Draw(image)
	c.HeapSortCB.Draw(image)
}
