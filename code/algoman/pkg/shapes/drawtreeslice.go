package shapes

import (
	"github.com/lei-cao/programming/code/v2/visualizer"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

type DrawTreeSlice struct {
	rs  *RectSlice
	dts []*DrawTree
}

func (d *DrawTreeSlice) Draw(egg *EbitenGG) {
	egg.img.Clear()
	egg.img.Fill(defaults.BackgroundColor)
	d.rs.Draw(egg.img)
}

func (d *DrawTreeSlice) Update(progress float64) {
	d.rs.Update(progress)
	if progress == 1 {
	}
}

func (d *DrawTreeSlice) NextStep(s visualizer.Stepper) {
	if step, ok := s.(*basicsort.Step); ok {
		if step.DoSwap() {
			dta := d.dts[step.A()]
			dtb := d.dts[step.B()]
			dta.DestX, dta.DestY = dtb.x, dtb.y
			dtb.DestX, dtb.DestY = dta.x, dta.y
			//dta.V, dtb.V = dtb.V, dta.V
		} else {
		}
		d.rs.NextStep(s)
	}
}
