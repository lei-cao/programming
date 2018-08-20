package ui

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"bytes"
	"github.com/lei-cao/programming/code/algoman/resources/images"
	"log"
)

var (
	playImg  *ebiten.Image
	pauseImg *ebiten.Image
)

func init() {
	play, _, err := image.Decode(bytes.NewReader(images.PLAY_ON_png))
	if err != nil {
		log.Fatal(err)
	}
	playImg, _ = ebiten.NewImageFromImage(play, ebiten.FilterDefault)

	pause, _, err := image.Decode(bytes.NewReader(images.PAUSE_ON_png))
	if err != nil {
		log.Fatal(err)
	}
	pauseImg, _ = ebiten.NewImageFromImage(pause, ebiten.FilterDefault)
}

func NewImageToggle(rect image.Rectangle) *ImageToggle {

	playToggle := &ImageToggle{
		Rect:   rect,
		ImgOff: playImg,
		ImgOn:  pauseImg,
	}
	return playToggle

}

type ImageToggle struct {
	ImgOn  *ebiten.Image
	ImgOff *ebiten.Image
	Rect   image.Rectangle
	On     bool

	mouseDown bool

	onPressed func(b *ImageToggle)
}

func (b *ImageToggle) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// @TODO Adding board border, make this better
		x -= 10
		y -= 10
		if b.Rect.Min.X <= x && x < b.Rect.Max.X && b.Rect.Min.Y <= y && y < b.Rect.Max.Y {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.On = !b.On
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *ImageToggle) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.23, 0.23)
	op.GeoM.Translate(float64(b.Rect.Min.X), float64(b.Rect.Min.Y))
	//if b.mouseDown {
	if b.On {
		dst.DrawImage(b.ImgOn, op)
	} else {
		dst.DrawImage(b.ImgOff, op)
	}
	//}
}

func (b *ImageToggle) SetOnPressed(f func(b *ImageToggle)) {
	b.onPressed = f
}
