// Copyright 2018 The Algoman Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
)

func NewGame() *Game {
	g := &Game{}
	g.autoPlay = true

	g.speed = defaults.Speed

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
	Board        *board.Board
	Screen       *ebiten.Image
	Controller   *Controller
	then         float64
	now          float64
	startTime    float64
	steps        visualizer.Stepper
	sorter       sorting.Sorter
	values       []int
	autoPlay     bool
	timing       func(progress float64) float64
	speed        int
	finished     bool
	finishedWait int
	algorithm    string
}

func (g *Game) Animate() error {
	g.now = makeTimestamp()
	var progress = (g.now - g.startTime) / float64(g.speed)
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
	if !(g.steps.Finished() && g.finishedWait > 5000) {
		g.Board.Draw()
	}

	g.Controller.Draw()

	g.Screen.Fill(color.Gray16{0xaaaa})
	g.Screen.DrawImage(g.Board.BoardImage, nil)
	g.Screen.DrawImage(g.Controller.Image, nil)
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
	if g.speed >= 100 {
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
	g.values = utils.Shuffle(30)

	g.Board = board.NewBoard(g.values)

	g.applyAlgorithm(g.algorithm)
}

func (g *Game) initController() {
	g.Controller = NewController()

	g.Controller.PlayToggle.On = true
	g.Controller.PlayToggle.SetOnPressed(func(b *ui.ToggleButton) {
		if b.On {
			g.Resume()
		} else {
			g.Stop()
		}
	})

	g.Controller.NextStepBtn.SetOnPressed(func(b *ui.Button) {
		g.NextStep()
	})

	g.Controller.SpeedUpBtn.SetOnPressed(func(b *ui.Button) {
		g.SpeedUp()
	})

	g.Controller.SpeedDownBtn.SetOnPressed(func(b *ui.Button) {
		g.SpeedDown()
	})

	g.Controller.RShuffleBtn.SetOnPressed(func(b *ui.Button) {
		g.Restart()
	})

	for _, ss := range g.Controller.SortSelect {
		ss.SetOnCheckChanged(func(c *ui.CheckBox) {
			if c.Checked() {
				for _, ss := range g.Controller.SortSelect {
					ss.UnCheck()
				}

				c.Check()
				g.algorithm = c.Value()
				g.Restart()
			}
		})

		if ss.Value() == "quick" {
			ss.Check()
			g.algorithm = ss.Value()
		}
	}
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
