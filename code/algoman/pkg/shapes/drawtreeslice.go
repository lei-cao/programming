package shapes

import (
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

type DrawTreeSlice struct {
	rs  *RectSlice
	tree *DrawTree
	dts []*DrawTree
	dtA *DrawTree
	dtB *DrawTree
}

func (d *DrawTreeSlice) Draw(img *ebiten.Image) {
	img.Clear()
	//ell.tree.Draw(b.egg, 0)
	//egg.dc.SetColor(defaults.BackgroundColor)
	//egg.dc.Clear()
	//egg.img.Fill(defaults.BackgroundColor)
	d.rs.Draw(img)
	d.tree.Draw(img, 0)
	//d.dts[0].Draw(egg, 0)
}

func (d *DrawTreeSlice) Update(progress float64) {
	d.rs.Update(progress)
	if d.dtA != nil && d.dtB != nil {
		d.dtA.color = color.Black
		d.dtA.Update(progress)
		d.dtB.color = defaults.ColorC
		d.dtB.Update(progress)
		if progress == 1 {
			d.dtA.V, d.dtB.V = d.dtB.V, d.dtA.V
			d.dtA = nil
			d.dtB = nil
		}

	}
}

func (d *DrawTreeSlice) NextStep(s visualizer.Stepper) {
	if step, ok := s.(*basicsort.Step); ok {
		if step.DoSwap() {
			d.dtA = d.dts[step.A()]
			d.dtB = d.dts[step.B()]
			d.dtA.DestX, d.dtA.DestY = d.dtB.x, d.dtB.y
			d.dtB.DestX, d.dtB.DestY = d.dtA.x, d.dtA.y
		} else {
		}
		d.rs.NextStep(s)
	}
}
