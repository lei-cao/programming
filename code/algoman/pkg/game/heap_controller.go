package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

func NewHeapController() Controllerable {
	c := new(HeapController)
	c.image, _ = ebiten.NewImage(defaults.ScreenWidth, defaults.ScreenHeight, ebiten.FilterDefault)
	return c
}

type HeapController struct {
	image *ebiten.Image
}

func (c *HeapController) Update() {

}
func (c *HeapController) Draw() {

}
func (c *HeapController) Image() *ebiten.Image {
	return c.image
}
