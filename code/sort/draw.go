package sort

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-canvas"
	"strconv"
)

var barWidth = 8
var barSpace = 2
var heightUnit = 5

func createCanvas(id string, length int) *js.Object {
	body := js.Global.Get("document").Call("getElementById", "body")
	obj := js.Global.Get("document").Call("createElement", "canvas")
	obj.Set("id", id)
	obj.Set("width", strconv.Itoa(canvasWidth(length)))
	obj.Set("height", strconv.Itoa(canvasHeight(length)))
	body.Call("appendChild", obj)
	return obj
}

func draw(nums []int, i int, j int, obj *js.Object) {

	can := canvas.New(obj)

	ctx := can.GetContext2D()

	ctx.ClearRect(0, 0, float64(canvasWidth(len(nums))), float64(canvasHeight(len(nums))))
	x := 0
	for k, v := range nums {
		ctx.FillStyle = "rgb(200, 0, 0)"
		if k == i {
			ctx.FillStyle = "green"
		}
		if k == j {
			ctx.FillStyle = "blue"
		}
		ctx.FillRect(float64(x+(barWidth + barSpace)*k), float64(canvasHeight(len(nums))-v*5), float64(barWidth), float64(v*heightUnit));
	}
}

func canvasWidth(length int) int {
	return barWidth * length + (length-1)*barSpace
}

func canvasHeight(length int) int {
	return length * heightUnit
}
