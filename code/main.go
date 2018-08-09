package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lei-cao/learning-cs-again/code/play"
	"github.com/lei-cao/learning-cs-again/code/v1/visualizer"
)

func main() {

	js.Global.Set("algorithm", map[string]interface{}{
		"Algorithm":        Algorithm,
		"Controller":       Controller,
		"ControllerConfig": ControllerConfig,
	})
}

func Algorithm() *js.Object {
	return js.MakeWrapper(new(visualizer.Visualizer))
}

func Controller() *js.Object {
	return js.MakeWrapper(new(controller.Controller))
}

func ControllerConfig() *js.Object {
	return js.MakeWrapper(new(controller.ControllerConfig))
}
