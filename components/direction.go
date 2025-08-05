package components

import "github.com/musztardem/zombic/vectors"

type Direction struct {
	X float64
	Y float64
}

func NewDirection(x, y float64) *Direction {
	return &Direction{
		X: x,
		Y: y,
	}
}

func NormalFromPositions(pos1, pos2 *Position) *Direction {
	dx := pos2.X - pos1.X
	dy := pos2.Y - pos1.Y

	nVecX, nVecY := vectors.Normal(dx, dy)

	return NewDirection(nVecX, nVecY)
}
