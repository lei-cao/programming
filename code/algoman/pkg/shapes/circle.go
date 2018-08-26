package shapes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/lei-cao/programming/code/algoman/resources/pngs"
	"github.com/lei-cao/programming/code/algoman/utils"
	"github.com/hajimehoshi/ebiten/text"
	"strconv"
	"github.com/lei-cao/programming/code/algoman/pkg/defaults"
	"golang.org/x/image/font"
)

func NewCircle(value int) *Circle {
	c := new(Circle)
	str := strconv.Itoa(value)
	c.Img = utils.ImgFromByte(pngs.FILLED_CIRCLE_48_png)

	bounds, _ := font.BoundString(utils.UiFont, str)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	width, height := c.Img.Size()
	x := (width - w) / 2
	y := height - (height - utils.UiFontMHeight) / 2

	text.Draw(c.Img, strconv.Itoa(value), utils.UiFont, x, y, defaults.BarColor)
	return c
}

type Circle struct {
	Img *ebiten.Image
}
