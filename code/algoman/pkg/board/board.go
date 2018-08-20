package board

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/pkg/shapes"
	"github.com/lei-cao/programming/code/algoman/pkg/consts"
	"github.com/lei-cao/programming/code/v2/visualizer"
)

func init() {
}

func NewBoard(values []int) *Board {
	b := &Board{values: values}
	b.rs = shapes.NewRectSlice(b.values)
	width, height := consts.ScreenWidth-consts.ScreenBorder*2, consts.ScreenHeight-consts.ScreenBorder*2
	b.BoardImage, _ = ebiten.NewImage(width, height, ebiten.FilterDefault)
	b.BoardImage.Fill(consts.BackgroundColor)
	return b
}

type Board struct {
	BoardImage *ebiten.Image
	Finished bool
	progress   float64
	values     []int
	rs         *shapes.RectSlice
}

func (b *Board) Draw() {
	b.BoardImage.Clear()
	b.BoardImage.Fill(consts.BackgroundColor)
	b.rs.Draw(b.BoardImage)
}

func (b *Board) Update(progress float64) {
	b.progress = progress
	b.rs.Update(progress)
}

func (b *Board) NextStep(step visualizer.Stepper) {
	b.rs.NextStep(step)
}

func (b *Board) Ready() bool {
	if b.progress == 1 {
		return true
	}
	return false
}
