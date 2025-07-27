package components

import (
	"github.com/musztardem/zombic/vectors"
)

type Path struct {
	Points []*Position
}

type PathFollowLoop struct {
	path                    *Path
	position                *Position
	velocity                *Velocity
	currentTargetPointIndex int
}

func NewPathFollowLoop(path *Path, velocity *Velocity) *PathFollowLoop {
	return &PathFollowLoop{
		path: path,
		position: &Position{
			X: path.Points[0].X,
			Y: path.Points[0].Y,
		},
		velocity:                velocity,
		currentTargetPointIndex: 0,
	}
}

func (pf *PathFollowLoop) Update() {
	pathPointsCount := len(pf.path.Points)
	targetPosition := pf.path.Points[pf.currentTargetPointIndex]

	dx := targetPosition.X - pf.position.X
	dy := targetPosition.Y - pf.position.Y

	nVecX, nVecY := vectors.Normal(dx, dy)

	pf.position.X += nVecX * pf.velocity.Val
	pf.position.Y += nVecY * pf.velocity.Val

	if pf.position.IsNear(targetPosition) {
		pf.currentTargetPointIndex = (pf.currentTargetPointIndex + 1) % pathPointsCount
	}
}

func (pf *PathFollowLoop) GetPosition() *Position {
	return pf.position
}
