package entities

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/config"
	"github.com/musztardem/zombic/images"
	"github.com/musztardem/zombic/vectors"
)

type Player struct {
	AnimatedSprite *components.AnimatedSprite
	Position       *components.Position
	Velocity       *components.Velocity
	Collider       *components.Collider
	Weapon         *Weapon
}

func NewPlayer(position *components.Position, velocity *components.Velocity) *Player {
	p := &Player{
		Position: position,
		Velocity: velocity,
		Weapon:   NewWeapon(position, 15),
	}
	p.loadAnimations()
	// Collider is dependent on the frame size so animation needs to be loaded first
	p.updateCollider()

	return p
}

func (p *Player) Update(enemies *[]EnemyBehaviour, missles *[]Missle) {
	p.handleMovement()
	p.Weapon.Update()
}

func (p *Player) handleMovement() {
	p.AnimatedSprite.Play("idle")

	var dx, dy float64

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.AnimatedSprite.Play("walk_up")
		dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.AnimatedSprite.Play("walk_down")
		dy = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.AnimatedSprite.Play("walk_left")
		dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.AnimatedSprite.Play("walk_right")
		dx = 1
	}

	nVecX, nVecY := vectors.Normal(dx, dy)
	p.Position.X += nVecX * p.Velocity.Val
	p.Position.Y += nVecY * p.Velocity.Val

	p.updateCollider()
	p.updateWeaponPosition()
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.AnimatedSprite.Get(), &opts)

	p.Weapon.Draw(p.Weapon.position, screen)

	if config.ColliderDebug {
		opt2 := ebiten.DrawImageOptions{}
		opt2.GeoM.Translate(p.Collider.Position.X, p.Collider.Position.Y)
		img := ebiten.NewImage(int(p.Collider.Width), int(p.Collider.Height))
		img.Fill(color.RGBA{255, 0, 0, 1})
		screen.DrawImage(img, &opt2)
	}
}

func (p *Player) updateCollider() {
	p.Collider = &components.Collider{
		Position: p.Position,
		Width:    float64(p.AnimatedSprite.Get().Bounds().Dx()),
		Height:   float64(p.AnimatedSprite.Get().Bounds().Dy()),
	}
}

func (p *Player) updateWeaponPosition() {
	p.Weapon.position = p.Position
}

func (p *Player) loadAnimations() {
	load := func(path string) *ebiten.Image {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatalf("failed to load image %v", err)
		}
		return img
	}

	// Load each unique frame only once
	frameRight1 := load("assets/sprites/Player/right1.png")
	frameRight2 := load("assets/sprites/Player/right2.png")
	frameRight3 := load("assets/sprites/Player/right3.png")
	frameUp1 := load("assets/sprites/Player/up1.png")
	frameUp2 := load("assets/sprites/Player/up2.png")
	frameUp3 := load("assets/sprites/Player/up3.png")
	frameDown1 := load("assets/sprites/Player/down1.png")
	frameDown2 := load("assets/sprites/Player/down2.png")
	frameDown3 := load("assets/sprites/Player/down3.png")

	// Mirror right frames for left
	frameLeft1 := images.Mirror(frameRight1)
	frameLeft2 := images.Mirror(frameRight2)
	frameLeft3 := images.Mirror(frameRight3)

	playerAnimatedSprite := components.NewAnimatedSprite()
	playerAnimatedSprite.RegisterAnimation(
		"walk_right",
		[]*ebiten.Image{
			frameRight1,
			frameRight2,
			frameRight1,
			frameRight3,
		},
		7,
	)
	playerAnimatedSprite.RegisterAnimation(
		"walk_left",
		[]*ebiten.Image{
			frameLeft1,
			frameLeft2,
			frameLeft1,
			frameLeft3,
		},
		7,
	)
	playerAnimatedSprite.RegisterAnimation(
		"walk_up",
		[]*ebiten.Image{
			frameUp1,
			frameUp2,
			frameUp1,
			frameUp3,
		},
		7,
	)
	playerAnimatedSprite.RegisterAnimation(
		"walk_down",
		[]*ebiten.Image{
			frameDown1,
			frameDown2,
			frameDown1,
			frameDown3,
		},
		7,
	)
	playerAnimatedSprite.RegisterAnimation(
		"idle",
		[]*ebiten.Image{
			frameRight1,
		},
		20,
	)

	p.AnimatedSprite = playerAnimatedSprite
}
