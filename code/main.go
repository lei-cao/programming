package main

import (
	"github.com/gopherjs/gopherjs/js"
	v2 "github.com/lei-cao/programming/code/v2/play"
	v1 "github.com/lei-cao/programming/code/v1/visualizer"
)

func main() {

	js.Global.Set("algorithm", map[string]interface{}{
		"Algorithm":        Algorithm,
		"Controller":       Controller,
		"ControllerConfig": ControllerConfig,
	})
}

func Algorithm() *js.Object {
	return js.MakeWrapper(new(v1.Visualizer))
}

func Controller() *js.Object {
	return js.MakeWrapper(new(v2.Controller))
}

func ControllerConfig() *js.Object {
	return js.MakeWrapper(new(v2.ControllerConfig))
}
