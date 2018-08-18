package main

import (
	"github.com/hajimehoshi/ebiten"
	"math"
	"github.com/lei-cao/programming/code/algoman/pkg/game"
	"github.com/lei-cao/programming/code/algoman/pkg/consts"
)

var (
	algoman *game.Game
)

func update(screen *ebiten.Image) error {
	algoman.Screen = screen
	return algoman.Animate()
}

func main() {
	algoman = game.NewGame()
	ebiten.Run(update, consts.ScreenWidth, consts.ScreenHeight, 1, "Hello world!")
}

func timing(progress float64) float64 {
	var x = 0.5
	return math.Pow(progress, 2) * ((x+1)*progress - x)
}
