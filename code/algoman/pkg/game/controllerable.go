package game

import "github.com/hajimehoshi/ebiten"

type Controllerable interface {
	Update()
	Draw()
	Image() *ebiten.Image
}
