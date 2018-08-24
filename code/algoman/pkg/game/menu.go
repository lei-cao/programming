package game

import (
	"github.com/lei-cao/programming/code/algoman/pkg/ui"
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"image"
)

var menuBtnGroup = &BtnGroup{
	x:         10 * defaults.DeviceScale,
	y:         defaults.ScreenHeight - 10*defaults.DeviceScale - defaults.ButtonMinHeight,
	margin:    10 * defaults.DeviceScale,
	height:    defaults.ButtonMinHeight,
	direction: btnGroupDirectionRight,
	rects:     []image.Rectangle{},
}

func NewMenu() *Menu {
	m := new(Menu)
	m.Image, _ = ebiten.NewImage(defaults.ScreenWidth, defaults.ScreenHeight, ebiten.FilterDefault)
	m.SortingButton = ui.NewButton(menuBtnGroup.nextRect(), "Sorting")
	menuBtnGroup.addRect(m.SortingButton.Rect)

	m.HeapButton = ui.NewButton(menuBtnGroup.nextRect(), "Heap")
	menuBtnGroup.addRect(m.SortingButton.Rect)
	return m
}

type Menu struct {
	Image *ebiten.Image

	SortingButton *ui.Button
	HeapButton    *ui.Button
}

func (m *Menu) Update(progress float64) {
	m.SortingButton.Update()
	m.HeapButton.Update()
}

func (m *Menu) Draw() {
	m.SortingButton.Draw(m.Image)
	m.HeapButton.Draw(m.Image)
}
