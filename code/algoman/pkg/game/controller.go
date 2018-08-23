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
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

var (
	controlX0 = 10 * defaults.DeviceScale
	controlY0 = 175 * defaults.DeviceScale
)

type btnGroupDirection int

const (
	btnGroupDirectionRight btnGroupDirection = iota
	btnGroupDirectionLeft
	btnGroupDirectionDown
	btnGroupDirectionUp
)

var gamePlayBtnGroup = &BtnGroup{
	x:         10 * defaults.DeviceScale,
	y:         controlY0,
	margin:    10 * defaults.DeviceScale,
	height:    defaults.ButtonMinHeight,
	direction: btnGroupDirectionRight,
	rects:     []image.Rectangle{},
}

type BtnGroup struct {
	x         int
	y         int
	margin    int
	height    int
	width     int
	direction btnGroupDirection
	rects     []image.Rectangle
}

func (i *BtnGroup) nextRect() image.Rectangle {
	var x0, y0, x1, y1 int
	if i.direction == btnGroupDirectionRight {
		x0 = i.x + i.width
		y0 = i.y
		x1 = x0
		y1 = y0 + i.height
	}
	return image.Rect(x0, y0, x1, y1)
}

func (i *BtnGroup) addRect(rect image.Rectangle) {
	i.width += rect.Dx() + i.margin
}

func NewController() *Controller {
	c := new(Controller)
	c.Image, _ = ebiten.NewImage(defaults.ScreenWidth, defaults.ScreenHeight, ebiten.FilterDefault)
	c.PlayToggle = ui.NewToggleButton(gamePlayBtnGroup.nextRect(), "Pause", "Play", false)
	gamePlayBtnGroup.addRect(c.PlayToggle.Rect)
	c.NextStepBtn = ui.NewButton(gamePlayBtnGroup.nextRect(), "Next")
	gamePlayBtnGroup.addRect(c.NextStepBtn.Rect)
	c.SpeedDownBtn = ui.NewButton(gamePlayBtnGroup.nextRect(), "-")
	gamePlayBtnGroup.addRect(c.SpeedDownBtn.Rect)
	c.SpeedUpBtn = ui.NewButton(gamePlayBtnGroup.nextRect(), "+")
	gamePlayBtnGroup.addRect(c.SpeedUpBtn.Rect)
	c.RShuffleBtn = ui.NewButton(gamePlayBtnGroup.nextRect(), "ReShuffle")
	gamePlayBtnGroup.addRect(c.RShuffleBtn.Rect)

	sorters := []struct {
		k string
		v string
	}{
		{
			"bubble", "Bubble Sort",
		},
		{
			"selection", "Selection Sort",
		},
		{
			"insertion", "Insertion Sort",
		},
		{
			"quick", "Quick Sort",
		},
		{
			"heap", "Heap Sort",
		},
	}

	for k, v := range sorters {
		cb := &ui.CheckBox{
			X:    10 * defaults.DeviceScale,
			Y:    gamePlayBtnGroup.nextRect().Max.Y + (20+35*k)*defaults.DeviceScale,
			Text: v.v,
		}
		cb.SetValue(v.k)
		c.SortSelect = append(c.SortSelect, cb)
	}

	return c
}

type Controller struct {
	Image *ebiten.Image

	PlayToggle   *ui.ToggleButton
	NextStepBtn  *ui.Button
	SpeedUpBtn   *ui.Button
	SpeedDownBtn *ui.Button
	RShuffleBtn  *ui.Button

	SortSelect []*ui.CheckBox
}

func (c *Controller) Update() {
	c.PlayToggle.Update()
	c.NextStepBtn.Update()
	c.SpeedDownBtn.Update()
	c.SpeedUpBtn.Update()
	c.RShuffleBtn.Update()

	for _, v := range c.SortSelect {
		v.Update()
	}
}

func (c *Controller) Draw() {
	c.PlayToggle.Draw(c.Image)
	c.NextStepBtn.Draw(c.Image)
	c.SpeedDownBtn.Draw(c.Image)
	c.SpeedUpBtn.Draw(c.Image)
	c.RShuffleBtn.Draw(c.Image)

	for _, v := range c.SortSelect {
		v.Draw(c.Image)
	}
}
