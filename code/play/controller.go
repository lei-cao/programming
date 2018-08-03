package controller

import (
	"github.com/lei-cao/learning-cs-again/code/sort"
	"github.com/gopherjs/gopherjs/js"
	"strconv"
	"github.com/oskca/gopherjs-canvas"
	"time"
	"runtime"
)

var barWidth = 8
var barSpace = 2
var heightUnit = 5

type Step struct {
	A int
	B int
}

type Controller struct {
	C           *canvas.Canvas
	Ctx         *canvas.Context2D
	Steps       []*Step
	CurrentStep int
	Config      *ControllerConfig
	StopChan    chan bool
}

var rectangles []*Rectangle
var moving bool

var defaultSize = 30

type ControllerConfig struct {
	Speed int `json:"speed"`
	Size  int `json:"size"`
}

func (c *ControllerConfig) SetSpeed(speed int) {
	c.Speed = speed
}

func (c *ControllerConfig) SetSize(size int) {
	c.Size = size
}

func (c *Controller) UpdateConfig(config *ControllerConfig) {
	c.Config = config
}

var StopChan chan bool

func (c *Controller) Init(id string, config *ControllerConfig) {
	c.Config = config
	c.Steps = []*Step{}
	c.StopChan = make(chan bool)
	if config.Size == 0 {
		c.Config.Size = defaultSize
	}

	rectangles = []*Rectangle{}
	var nums = sort.Shuffle(c.Config.Size)

	for k, v := range nums {
		r := NewRect(c.Config.Size, k, v)
		rectangles = append(rectangles, r)
	}

	for i := 0; i < c.Config.Size; i++ {
		for j := 0; j < c.Config.Size-1; j++ {
			if nums[j] > nums[j+1] {
				step := &Step{}
				nums[j], nums[j+1] = nums[j+1], nums[j]
				step.A = j
				step.B = j + 1
				c.Steps = append(c.Steps, step)
			}
		}
	}

	obj := createCanvas(id, c.Config.Size)
	c.C = canvas.New(obj)
	c.Ctx = c.C.GetContext2D()
	c.animate()
}

func (c *Controller) Run() {
	go func() {
		for {
			select {
			case <-c.StopChan:
				runtime.Goexit()
				break
			case <-time.After(time.Duration(c.Config.Speed) * time.Millisecond):
				if c.CurrentStep == len(c.Steps) {
					runtime.Goexit()
					break
				}
				v := c.Steps[c.CurrentStep]
				rectangles[v.A].ToIndex = v.B
				rectangles[v.B].ToIndex = v.A

				rectangles[v.A], rectangles[v.B] = rectangles[v.B], rectangles[v.A]
				c.CurrentStep++
				if rectangles[v.A].Moving || rectangles[v.B].Moving {
					continue
				}
			}
		}
	}()
}

func (c *Controller) Stop() {
	go func() {
		c.StopChan <- true
	}()
}

type Rectangle struct {
	Index   int
	ToIndex int
	Total   int
	Value   int
	Left    float64
	Top     float64
	Width   float64
	Height  float64
	Moving  bool
}

func (r *Rectangle) Draw(ctx *canvas.Context2D) {
	ctx.FillStyle = "#B9314F"
	if r.Index != r.ToIndex {
		ctx.FillStyle = "#12355B"
	}
	ctx.FillRect(r.Left, r.Top, r.Width, r.Height)
}

func (r *Rectangle) Update(ctx *canvas.Context2D) {
	if r.Left != r.ToLeft() {
		r.Moving = true
		if r.Index > r.ToIndex {
			r.Left -= 2
		} else if r.Index < r.ToIndex {
			r.Left += 2
		}
	} else {
		r.Index = r.ToIndex
		r.Moving = false
	}

	r.Draw(ctx)
}

func (r *Rectangle) ToLeft() float64 {
	return float64((barWidth + barSpace) * r.ToIndex)
}

func NewRect(total int, index int, value int) *Rectangle {
	r := new(Rectangle)
	r.Index = index
	r.ToIndex = index
	r.Total = total
	r.Value = value
	r.Left = float64((barWidth + barSpace) * index)
	r.Top = float64(canvasHeight(total) - value*heightUnit)
	r.Width = float64(barWidth)
	r.Height = float64(value * heightUnit)
	return r
}

func (c *Controller) animate() {
	js.Global.Call("requestAnimationFrame", c.animate)

	c.Ctx.ClearRect(0, 0, float64(canvasWidth(c.Config.Size)), float64(canvasHeight(c.Config.Size)))
	for _, v := range rectangles {
		v.Update(c.Ctx)
	}
}

func createCanvas(id string, size int) *js.Object {
	body := js.Global.Get("document").Call("getElementById", id)
	obj := js.Global.Get("document").Call("createElement", "canvas")
	obj.Set("width", strconv.Itoa(canvasWidth(size)))
	obj.Set("height", strconv.Itoa(canvasHeight(size)))
	body.Set("innerHTML", "")
	body.Call("appendChild", obj)
	return obj
}

func canvasWidth(size int) int {
	return barWidth*size + (size-1)*barSpace
}

func canvasHeight(size int) int {
	return size * heightUnit
}
