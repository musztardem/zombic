package components

import "math"

type Position struct {
	X float64
	Y float64
}

func (p *Position) IsNear(p2 *Position) bool {
	dx := p.X - p2.X
	dy := p.Y - p2.Y

	return math.Round(dx) == 0 && math.Round(dy) == 0
}

func (p *Position) Translate(dx, dy float64) *Position {
	return &Position{
		X: p.X + dx,
		Y: p.Y + dy,
	}
}

func (p *Position) DistanceTo(p2 *Position) float64 {
	dx := p2.X - p.X
	dy := p2.Y - p.Y

	return math.Sqrt(dx*dx + dy*dy)
}
