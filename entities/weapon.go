package entities

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/musztardem/zombic/components"
)

type Shootable interface {
	ShootAt(targetPosition *components.Position) *Missle
}

type Weapon struct {
	sprite               *ebiten.Image
	position             *components.Position
	shootingSpeed        int // ticks interval == lesser is faster
	shootingSpeedCounter int
}

var (
	weaponImageLoaded  sync.Once
	weaponLoadedSprite *ebiten.Image
)

func NewWeapon(position *components.Position, shootingSpeed int) *Weapon {
	weapon := &Weapon{
		shootingSpeed: shootingSpeed,
	}
	weapon.loadImage()

	return weapon
}

func (w *Weapon) CanShoot() bool {
	return w.shootingSpeed <= w.shootingSpeedCounter
}

func (w *Weapon) ShootAt(targetPosition *components.Position) *Missle {
	if w.CanShoot() {
		w.shootingSpeedCounter = 0
	} else {
		return nil
	}

	direction := components.NormalFromPositions(w.BarrelPosition(), targetPosition)

	return NewMissle(w.BarrelPosition(), direction)
}

func (w *Weapon) BarrelPosition() *components.Position {
	rect := w.sprite.Bounds()
	rectMax := rect.Max

	posX := float64(rectMax.X)
	posY := float64(rectMax.Y) - (0.5 * float64(rect.Dy()))

	return &components.Position{
		X: w.position.X + posX,
		Y: w.position.Y + posY,
	}
}

func (w *Weapon) Draw(playerPosition *components.Position, screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(playerPosition.X+1, playerPosition.Y+7)
	screen.DrawImage(w.sprite, opts)
}

func (w *Weapon) Update() {
	w.shootingSpeedCounter++
}

func (w *Weapon) loadImage() {
	w.loadImageOnce()
}

func (w *Weapon) loadImageOnce() {
	weaponImageLoaded.Do(func() {
		weaponImage, _, err := ebitenutil.NewImageFromFile("assets/sprites/Pickable/shotgun.png")
		if err != nil {
			log.Fatalf("failed to load image %v", err)
		}

		weaponLoadedSprite = weaponImage
	})

	w.sprite = weaponLoadedSprite
}
