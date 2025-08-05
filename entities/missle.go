package entities

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/config"
)

type Missle struct {
	sprite      *ebiten.Image
	position    *components.Position
	direction   *components.Direction
	speed       float64
	Collider    *components.Collider
	IsRemovable bool
}

var (
	missleImageLoaded  sync.Once
	missleLoadedSprite *ebiten.Image
)

func NewMissle(position, targetPosition *components.Position) *Missle {
	m := &Missle{
		position: &components.Position{
			X: position.X,
			Y: position.Y,
		},
		direction: components.NormalFromPositions(position, targetPosition),
		speed:     5.0,
	}
	m.loadSprite()
	m.updateCollider()

	return m
}

func (m *Missle) Update() {
	m.position.X = m.position.X + m.direction.X*m.speed
	m.position.Y = m.position.Y + m.direction.Y*m.speed

	m.updateCollider()
}

func (m *Missle) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, &opts)

	if config.ColliderDebug {
		opt2 := ebiten.DrawImageOptions{}
		opt2.GeoM.Translate(m.Collider.Position.X, m.Collider.Position.Y)
		img := ebiten.NewImage(int(m.Collider.Width), int(m.Collider.Height))
		img.Fill(color.RGBA{255, 0, 0, 1})
		screen.DrawImage(img, &opt2)
	}

}

func (m *Missle) updateCollider() {
	m.Collider = &components.Collider{
		Position: m.position,
		Width:    float64(m.sprite.Bounds().Dx()),
		Height:   float64(m.sprite.Bounds().Dy()),
	}
}

func (m *Missle) loadSprite() {
	missleImageLoaded.Do(func() {
		missleLoadedSprite = ebiten.NewImage(2, 2)
		missleLoadedSprite.Fill(color.RGBA{120, 0, 0, 255})
	})

	m.sprite = missleLoadedSprite
}
