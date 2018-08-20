package ui

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
)

func NewImageButton(rect image.Rectangle, imgOn *ebiten.Image, imgOff *ebiten.Image) *ImageButton {

	stopBtn := &ImageButton{
		Rect:   rect,
		ImgOff: imgOff,
		ImgOn:  imgOn,
	}
	return stopBtn

}

type ImageButton struct {
	ImgOn  *ebiten.Image
	ImgOff *ebiten.Image
	Rect   image.Rectangle

	mouseDown bool

	onPressed func(b *ImageButton)
}

func (b *ImageButton) Update() {
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
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *ImageButton) Draw(dst *ebiten.Image) {
	t := imageTypeButton
	if b.mouseDown {
		t = imageTypeButtonPressed
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.23, 0.23)
	op.GeoM.Translate(float64(b.Rect.Min.X), float64(b.Rect.Min.Y))
	if t == imageTypeButtonPressed {
		dst.DrawImage(b.ImgOn, op)
	} else {
		dst.DrawImage(b.ImgOff, op)
	}
}

func (b *ImageButton) SetOnPressed(f func(b *ImageButton)) {
	b.onPressed = f
}
