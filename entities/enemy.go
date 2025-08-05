package entities

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/config"
	"github.com/musztardem/zombic/vectors"
)

type EnemyState int

const (
	StateAttacking EnemyState = iota
	StateHit
	StateDead
)

type EnemyBehaviour interface {
	Update([]*components.Collider) error
	Draw(screen *ebiten.Image)
	GetPosition() *components.Position
	GetAnimation() *ebiten.Image
	GetCollider() *components.Collider

	// State management methods
	MarkAsHit()
	MarkAsDead()
	IsDead() bool
}

type Enemy struct {
	AnimatedSprite *components.AnimatedSprite
	Position       *components.Position
	Velocity       *components.Velocity
	TargetPosition *components.Position
	Collider       *components.Collider
	isDead         bool
	State          EnemyState
	hitCooldown    int
}

func (e *Enemy) Update(colliders []*components.Collider) error {
	e.AnimatedSprite.Play("idle")
	dx := e.TargetPosition.X - e.Position.X
	dy := e.TargetPosition.Y - e.Position.Y
	nVecX, nVecY := vectors.Normal(dx, dy)

	e.handleMoveWithCollisions(nVecX, nVecY, colliders)

	switch e.State {
	case StateAttacking:
		if nVecX > nVecY && nVecX > 0 {
			e.AnimatedSprite.Play("walk_right")
		} else if nVecX < nVecY && nVecX < 0 {
			e.AnimatedSprite.Play("walk_left")
		} else if nVecY > nVecX && nVecY > 0 {
			e.AnimatedSprite.Play("walk_down")
		} else if nVecY < nVecX && nVecY < 0 {
			e.AnimatedSprite.Play("walk_up")
		}

	case StateHit:
		if nVecX > nVecY && nVecX > 0 {
			e.AnimatedSprite.Play("damaged_right")
		} else if nVecX < nVecY && nVecX < 0 {
			e.AnimatedSprite.Play("damaged_left")
		} else if nVecY > nVecX && nVecY > 0 {
			e.AnimatedSprite.Play("damaged_down")
		} else if nVecY < nVecX && nVecY < 0 {
			e.AnimatedSprite.Play("damaged_up")
		}

		e.hitCooldown--
		if e.hitCooldown <= 0 {
			e.State = StateAttacking
			e.Velocity.Val = e.Velocity.Val * 3
		}
	}

	e.updateCollider()

	return nil
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.Position.X, e.Position.Y)
	screen.DrawImage(e.AnimatedSprite.Get(), opts)

	if config.ColliderDebug {
		opt2 := ebiten.DrawImageOptions{}
		opt2.GeoM.Translate(e.Collider.Position.X, e.Collider.Position.Y)
		img := ebiten.NewImage(int(e.Collider.Width), int(e.Collider.Height))
		img.Fill(color.RGBA{255, 0, 0, 1})
		screen.DrawImage(img, &opt2)
	}
}

func (e *Enemy) GetPosition() *components.Position {
	return e.Position
}

func (e *Enemy) GetAnimation() *ebiten.Image {
	return e.AnimatedSprite.Get()
}

func (e *Enemy) GetCollider() *components.Collider {
	return e.Collider
}

func (e *Enemy) MarkAsHit() {
	// we want to slow down the enemy on the state transition
	if e.State != StateHit {
		e.Velocity.Val = e.Velocity.Val / 3
	}

	e.State = StateHit
	e.hitCooldown = 30
}

func (e *Enemy) MarkAsDead() {
	e.State = StateDead
}

func (e *Enemy) IsDead() bool {
	return e.State == StateDead
}

func (e *Enemy) updateCollider() {
	offsetX := 2.0
	offsetY := 2.0

	e.Collider = &components.Collider{
		Position: e.Position.Translate(offsetX, offsetY),
		Width:    float64(e.AnimatedSprite.Get().Bounds().Dx()) - 2*offsetX,
		Height:   float64(e.AnimatedSprite.Get().Bounds().Dy()) - 2*offsetY,
	}
}

func (e *Enemy) handleMoveWithCollisions(vecX, vecY float64, colliders []*components.Collider) {
	e.handleYMoveWithCollisions(vecY, colliders)
	e.handleXMoveWithCollisions(vecX, colliders)
}

func (e *Enemy) handleXMoveWithCollisions(vecX float64, colliders []*components.Collider) {
	nextX := e.Position.X + vecX*e.Velocity.Val

	isNextMoveLeft := nextX < e.Position.X
	isNextMoveRight := nextX > e.Position.X

	canMoveX := true
	nextColliderX := &components.Collider{
		Position: &components.Position{X: nextX, Y: e.Position.Y},
		Width:    e.Collider.Width,
		Height:   e.Collider.Height,
	}

	for _, other := range colliders {
		if e.Collider == other {
			continue
		}
		if nextColliderX.CollidesWith(other) && nextColliderX.CollidesFromRightWith(other) && isNextMoveRight {
			canMoveX = false
			break
		}
		if nextColliderX.CollidesWith(other) && nextColliderX.CollidesFromLeftWith(other) && isNextMoveLeft {
			canMoveX = false
			break
		}
	}

	if canMoveX {
		e.Position.X = nextX
	}
}

func (e *Enemy) handleYMoveWithCollisions(vecY float64, colliders []*components.Collider) {
	nextY := e.Position.Y + vecY*e.Velocity.Val

	isNextMoveUp := nextY > e.Position.Y
	isNextMoveDown := nextY < e.Position.Y

	canMoveY := true
	nextColliderY := &components.Collider{
		Position: &components.Position{X: e.Position.X, Y: nextY},
		Width:    e.Collider.Width,
		Height:   e.Collider.Height,
	}
	for _, other := range colliders {
		if e.Collider == other {
			continue
		}
		if nextColliderY.CollidesWith(other) && nextColliderY.CollidesFromDownWith(other) && isNextMoveDown {
			canMoveY = false
			break
		}
		if nextColliderY.CollidesWith(other) && nextColliderY.CollidesFromTopWith(other) && isNextMoveUp {
			canMoveY = false
			break
		}
	}

	if canMoveY {
		e.Position.Y = nextY
	}
}
