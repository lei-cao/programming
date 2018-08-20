package game

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/lei-cao/programming/code/algoman/pkg/board"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting/basicsort"
	"github.com/lei-cao/programming/code/v2/algorithms/sorting"
	"github.com/lei-cao/programming/code/utils"
	"github.com/lei-cao/programming/code/v2/visualizer"
	"time"
	"github.com/lei-cao/programming/code/algoman/pkg/ui"
	"math"
	"github.com/lei-cao/programming/code/algoman/pkg/consts"
)

func NewGame() *Game {
	g := &Game{}
	g.autoPlay = true

	g.speed = 300

	g.steps = visualizer.NewFirstStep()

	g.timing = func(progress float64) float64 {
		var x = 0.5
		return math.Pow(progress, 2) * ((x+1)*progress - x)
	}

	g.initController()
	g.initAlgorithm()

	return g
}

// Game represents a game state.
type Game struct {
	Board      *board.Board
	Screen     *ebiten.Image
	Controller *Controller
	then       float64
	now        float64
	startTime  float64
	steps      visualizer.Stepper
	sorter     sorting.Sorter
	values     []int
	autoPlay   bool
	timing     func(progress float64) float64
	speed      float64
	finished   bool
	finishedWait int
}

func (g *Game) Animate() error {
	g.now = makeTimestamp()
	var progress = (g.now - g.startTime) / g.speed
	if progress > 1 {
		progress = 1
	}

	g.then = g.now
	err := g.Update(g.timing(progress))
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
	g.Controller.Update()
	if progress == 1 {
		if g.autoPlay {
			g.NextStep()
		} else {
			g.Stop()
		}
	}

	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw() {
	if !(g.steps.Finished() && g.finishedWait > 100) {
		g.Board.Draw()
	}

	g.Controller.Draw(g.Board.BoardImage)

	g.Screen.Fill(color.Gray16{0xaaaa})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(consts.ScreenBorder), float64(consts.ScreenBorder))
	g.Screen.DrawImage(g.Board.BoardImage, op)
}

func (g *Game) NextStep() {
	if !g.Board.Ready() {
		return
	}
	if g.steps.Finished() {
		g.finishedWait ++
		return
	}

	g.startTime = makeTimestamp()
	g.Board.NextStep(g.steps.NextStep())
}

func (g *Game) Stop() {
	g.autoPlay = false
}

func (g *Game) Resume() {
	g.autoPlay = true
}

func (g *Game) SpeedUp() {
	if g.speed >= 100 && g.speed < 2000 {
		g.speed -= 100
	}
	if g.speed <= 0 {
		g.speed = 10
	}
}

func (g *Game) SpeedDown() {
	if g.speed < 2000 {
		g.speed += 100
	}
}

func (g *Game) Restart() {
	g.initAlgorithm()
}

func (g *Game) initAlgorithm() {
	g.values = utils.Shuffle(35)

	g.Board = board.NewBoard(g.values)

	g.applyAlgorithm("quick")
}

func (g *Game) initController() {
	g.Controller = NewController()

	g.Controller.PlayToggle.On = true
	g.Controller.PlayToggle.SetOnPressed(func(b *ui.ImageToggle) {
		if b.On {
			g.Resume()
		} else {
			g.Stop()
		}
	})

	g.Controller.NextStepBtn.SetOnPressed(func(b *ui.ImageButton) {
		g.NextStep()
	})

	g.Controller.SpeedUpBtn.SetOnPressed(func(b *ui.ImageButton) {
		g.SpeedUp()
	})

	g.Controller.SpeedDownBtn.SetOnPressed(func(b *ui.ImageButton) {
		g.SpeedDown()
	})

	g.Controller.RestartBtn.SetOnPressed(func(b *ui.ImageButton) {
		g.Restart()
	})
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
