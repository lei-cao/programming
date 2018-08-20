package main

import (
	"github.com/hajimehoshi/ebiten"
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
	ebiten.Run(update, consts.ScreenWidth, consts.ScreenHeight, 1.1, "Algoman")
}
