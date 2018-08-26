package utils

import (
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gomono"
	"log"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"github.com/hajimehoshi/ebiten/text"
	"image"
	"image/color"
	"github.com/hajimehoshi/ebiten"
)

var (
	UiFont        font.Face
	UiFontMHeight int
)

func init() {

	tt, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}
	UiFont = truetype.NewFace(tt, &truetype.Options{
		Size:    float64(12 * defaults.DeviceScale),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	b, _, _ := UiFont.GlyphBounds('M')
	UiFontMHeight = (b.Max.Y - b.Min.Y).Ceil()
}

func DrawString(dst *ebiten.Image, str string, rect image.Rectangle, color color.Color) {
	bounds, _ := font.BoundString(UiFont, str)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := rect.Min.X + (rect.Dx()-w)/2
	y := rect.Max.Y - (rect.Dy()-UiFontMHeight)/2
	text.Draw(dst, str, UiFont, x, y, color)
}
