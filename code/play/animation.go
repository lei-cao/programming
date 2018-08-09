package play

import (
	"github.com/lei-cao/learning-cs-again/code/visualizer"
	"github.com/gopherjs/gopherjs/js"
	"math"
	"time"
)

type Animator interface {
	StartAnimating()
	Animate(timestamp float64)
	Draw(progress float64)
	UpdateDuration(duration float64)
	SetScreen(screener visualizer.Screener)
	AddStep(a, b int, doSwap bool)
	Stop()
	Resume()
	NextStep()
}

func NewAnimation() Animator {
	a := new(Animation)
	a.steps = visualizer.NewStep()
	a.autoUpdate = true
	a.fps = 60
	a.fpdInterval = 1000 / float64(a.fps)
	return a
}

type Animation struct {
	steps          visualizer.Stepper
	screen         visualizer.Screener
	animating      bool
	autoUpdate     bool
	fps            int
	fpdInterval    float64
	startTime      float64
	now            float64
	then           float64
	elapsed        float64
	duration       float64
	timing         func(progress float64) float64
	animationFrame *js.Object
}

func (a *Animation) AddStep(ia, ib int, doSwap bool) {
	a.steps.AddStep(ia, ib, doSwap)
}

func (a *Animation) UpdateDuration(duration float64) {
	a.duration = duration
}

func (a *Animation) SetScreen(screen visualizer.Screener) {
	a.screen = screen
}

func (a *Animation) StartAnimating() {
	a.then = makeTimestamp()
	a.startTime = a.then

	a.timing = func(progress float64) float64 {
		var x = 0.5
		return math.Pow(progress, 2) * ((x+1)*progress - x)
	}
	a.animationFrame = js.Global.Call("requestAnimationFrame", a.Animate)
}

func (a *Animation) Animate(timestamp float64) {
	a.animationFrame = js.Global.Call("requestAnimationFrame", a.Animate)
	a.now = makeTimestamp()
	a.elapsed = a.now - a.then
	a.animating = true

	if a.elapsed > a.fpdInterval {

		var progress = float64(a.now-a.startTime) / a.duration
		if progress > 1 {
			progress = 1
		}
		a.then = a.now - math.Mod(a.elapsed, a.fpdInterval)

		a.screen.Clear()
		a.Draw(a.timing(progress))
	}
}

// Draw the screen
func (a *Animation) Draw(progress float64) {
	a.screen.Draw(progress)
	if a.screen.Ready() {
		if a.autoUpdate {
			a.NextStep()
		} else {
			// Not auto update, cancelAnimationFrame
			a.Stop()
		}
	}
}

// Stop auto running. Switch to manual control
func (a *Animation) Stop() {
	js.Global.Call("cancelAnimationFrame", a.animationFrame)
	a.animating = false
	a.autoUpdate = false
}

// Resume auto running
func (a *Animation) Resume() {
	if !a.autoUpdate {
		a.autoUpdate = true
		a.StartAnimating()
	}
}

// Do next step from the steps queue
// Being triggered manually or automatically
func (a *Animation) NextStep() {
	if a.steps.Finished() {
		js.Global.Call("cancelAnimationFrame", a.animationFrame)
		return
	}

	if !a.animating {
		a.StartAnimating()
	}
	if !a.screen.Ready() {
		return
	}
	a.startTime = makeTimestamp()
	a.screen.Update(a.steps.NextStep())
}

func makeTimestamp() float64 {
	return float64(time.Now().UnixNano()) / float64(time.Millisecond)
}
