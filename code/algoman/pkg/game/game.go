package game

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/lei-cao/programming/code/algoman/pkg/consts"
	"github.com/lei-cao/programming/code/algoman/pkg/board"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting"
	"github.com/lei-cao/programming/code/utils"
	"github.com/lei-cao/programming/code/v2/visualizer"
	"time"
)

func NewGame() *Game {
	g := &Game{}
	g.autoPlay = true

	g.values = utils.Shuffle(30)

	g.Board = board.NewBoard(g.values)

	g.steps = visualizer.NewFirstStep()

	g.applyAlgorithm("heap")

	return g
}

// Game represents a game state.
type Game struct {
	Board     *board.Board
	Screen    *ebiten.Image
	then      float64
	now       float64
	startTime float64
	steps     visualizer.Stepper
	sorter    sorting.Sorter
	values    []int
	autoPlay  bool
}

func (g *Game) Animate() error {
	g.now = makeTimestamp()
	var progress = (g.now - g.startTime) / 1000
	if progress > 1 {
		progress = 1
	}

	g.then = g.now
	err := g.Update(progress)
	if err != nil {
		return err
	}
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	g.Draw()
	return nil
}

func (g *Game) Update(progress float64) error {
	g.Board.Update(progress)
	if progress == 1 {
		if g.autoPlay {
			g.NextStep()
		} else {
			g.Stop()
		}
	}

	return nil
}

func (g *Game) NextStep() {
	g.Board.NextStep(g.steps.NextStep())
	g.startTime = makeTimestamp()
}

// Draw draws the current game to the given screen.
func (g *Game) Draw() {
	g.Board.Draw()
	g.Screen.Fill(color.Gray16{0xaaaa})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(consts.ScreenBorder), float64(consts.ScreenBorder))
	g.Screen.DrawImage(g.Board.BoardImage, op)
}

func (g *Game) Stop() {

}

func (g *Game) applyAlgorithm(id string) {
	switch id {
	case "bubble":
		g.sorter = basicsort.NewBubbleSort()
	case "selection":
		g.sorter = basicsort.NewSelectionSort()
	case "insertion":
		g.sorter = basicsort.NewInsertionSort()
	case "quick":
		g.sorter = basicsort.NewQuickSort()
	case "heap":
		g.sorter = basicsort.NewHeapSort()
		//case "topDownMergeSort":
		//	s := merge.NewScreen(c.config.Id, c.config.Size, c.nums)
		//	c.animation.SetScreen(s)
		//	c.sorter = mergesort.NewTopDownMergeSort()
	}
	g.sorter.Sort(g.values)
	g.steps = g.sorter.Steps()
}

func makeTimestamp() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Millisecond)
}
