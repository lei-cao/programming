package ui

type Point struct {
	X float64
	Y float64
}

func (p *Point) MoveTo(dest Point, progress float64) {
	p.X += (dest.X - p.X) * progress
	p.Y += (dest.Y - p.Y) * progress
}

func (p *Point) Equals(b Point) bool {
	return p.X == b.X && p.Y == b.Y
}