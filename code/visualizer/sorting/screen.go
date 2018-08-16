package sorting

import (
	"github.com/oskca/gopherjs-canvas"
	"github.com/lei-cao/programming/code/visualizer"
	"github.com/lei-cao/programming/code/visualizer/defaults"
	"github.com/lei-cao/programming/code/visualizer/ui"
)

// The screen including the elements on the canvas
// Maintain the ready for next state
type BaseScreen struct {
	Id      string
	Size    int
	C       *canvas.Canvas
	Ctx     *canvas.Context2D
	Element ui.Elementer
}

func (s *BaseScreen) Ready() bool {
	return s.Element.Ready()
}

func (s *BaseScreen) Clear() {
	s.Ctx.ClearRect(0, 0, s.Width(), s.Height())
}

func (s *BaseScreen) Width() float64 {
	return s.Element.Width()
}

func (s *BaseScreen) Height() float64 {
	return s.Element.Height()
}

func (s *BaseScreen) Draw(progress float64) {
	s.Ctx.FillStyle = defaults.DefaultColor.BackgroundColor
	s.Ctx.FillRect(0, 0, s.Width(), s.Height())

	s.Element.Draw(progress)
}

func (s *BaseScreen) Update(i visualizer.Stepper) {
	s.Element.Update(i)
}
