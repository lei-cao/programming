package board

import (
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/hajimehoshi/ebiten"
)

type Boarder interface {
	Draw()
	Update(progress float64)
	NextStep(step visualizer.Stepper)
	Ready() bool
	Image() *ebiten.Image
	Steps(id string) visualizer.Stepper
}
