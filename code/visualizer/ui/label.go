package ui

import "github.com/oskca/gopherjs-canvas"

func NewLabel(text string, color string, font string, ctx *canvas.Context2D, startPoint *Point, maxWidth float64) *Label {
	l := &Label{
		Text:       text,
		Color:      color,
		Font:       font,
		Ctx:        ctx,
		StartPoint: startPoint,
		MaxWidth:   maxWidth,
	}
	return l
}

type Label struct {
	Text       string
	Color      string
	Font       string
	Ctx        *canvas.Context2D
	StartPoint *Point
	MaxWidth   float64
}

func (l *Label) Draw() {
	l.Ctx.FillStyle = l.Color
	l.Ctx.Font = l.Font
	l.Ctx.FillText(l.Text, l.StartPoint.X, l.StartPoint.Y, l.MaxWidth);
}
