package entities

import (
	"log"
	"math/rand/v2"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/images"
)

type KidZombie struct {
	*Enemy
}

const (
	KZ_VELOCITY_LOWER_BOUND  = 0.5
	KZ_VELOCITY_HIGHER_BOUND = 1.0
)

var (
	kidZombieFramesLoaded                       sync.Once
	kzFrameLeft1, kzFrameLeft2, kzFrameLeft3    *ebiten.Image
	kzFrameUp1, kzFrameUp2, kzFrameUp3          *ebiten.Image
	kzFrameDown1, kzFrameDown2, kzFrameDown3    *ebiten.Image
	kzFrameRight1, kzFrameRight2, kzFrameRight3 *ebiten.Image

	kzFrameDamagedLeft1, kzFrameDamagedLeft2, kzFrameDamagedLeft3    *ebiten.Image
	kzFrameDamagedUp1, kzFrameDamagedUp2, kzFrameDamagedUp3          *ebiten.Image
	kzFrameDamagedDown1, kzFrameDamagedDown2, kzFrameDamagedDown3    *ebiten.Image
	kzFrameDamagedRight1, kzFrameDamagedRight2, kzFrameDamagedRight3 *ebiten.Image
)

func NewKidZombie(position, targetPosition *components.Position) *KidZombie {
	kidZombie := &KidZombie{
		Enemy: &Enemy{
			Position:       position,
			TargetPosition: targetPosition,
			Velocity:       generateKidZombieVelicity(),
		},
	}

	kidZombie.loadAnimations()
	kidZombie.updateCollider()

	return kidZombie
}

func (kz *KidZombie) loadAnimations() {
	kz.loadEnemyFramesOnce()

	enemyAnimatedSprite := components.NewAnimatedSprite()
	enemyAnimatedSprite.RegisterAnimation(
		"walk_right",
		[]*ebiten.Image{
			kzFrameRight1,
			kzFrameRight2,
			kzFrameRight1,
			kzFrameRight3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_left",
		[]*ebiten.Image{
			kzFrameLeft1,
			kzFrameLeft2,
			kzFrameLeft1,
			kzFrameLeft3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_up",
		[]*ebiten.Image{
			kzFrameUp1,
			kzFrameUp2,
			kzFrameUp1,
			kzFrameUp3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_down",
		[]*ebiten.Image{
			kzFrameDown1,
			kzFrameDown2,
			kzFrameDown1,
			kzFrameDown3,
		},
		7,
	)

	enemyAnimatedSprite.RegisterAnimation(
		"damaged_right",
		[]*ebiten.Image{
			kzFrameDamagedRight1,
			kzFrameDamagedRight2,
			kzFrameDamagedRight3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_left",
		[]*ebiten.Image{
			kzFrameDamagedLeft1,
			kzFrameDamagedLeft2,
			kzFrameDamagedLeft3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_up",
		[]*ebiten.Image{
			kzFrameDamagedUp1,
			kzFrameDamagedUp2,
			kzFrameDamagedUp3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_down",
		[]*ebiten.Image{
			kzFrameDamagedDown1,
			kzFrameDamagedDown2,
			kzFrameDamagedDown3,
		},
		3,
	)

	enemyAnimatedSprite.RegisterAnimation(
		"idle",
		[]*ebiten.Image{
			kzFrameDown1,
		},
		20,
	)

	kz.AnimatedSprite = enemyAnimatedSprite
}

func (kz *KidZombie) loadEnemyFramesOnce() {
	kidZombieFramesLoaded.Do(func() {
		load := func(path string) *ebiten.Image {
			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				log.Fatalf("failed to load image %v", err)
			}
			return img
		}

		// Walking
		kzFrameLeft1 = load("assets/sprites/KidZombie/left1.png")
		kzFrameLeft2 = load("assets/sprites/KidZombie/left2.png")
		kzFrameLeft3 = load("assets/sprites/KidZombie/left3.png")
		kzFrameUp1 = load("assets/sprites/KidZombie/up1.png")
		kzFrameUp2 = load("assets/sprites/KidZombie/up2.png")
		kzFrameUp3 = load("assets/sprites/KidZombie/up3.png")
		kzFrameDown1 = load("assets/sprites/KidZombie/down1.png")
		kzFrameDown2 = load("assets/sprites/KidZombie/down2.png")
		kzFrameDown3 = load("assets/sprites/KidZombie/down3.png")
		kzFrameRight1 = images.Mirror(kzFrameLeft1)
		kzFrameRight2 = images.Mirror(kzFrameLeft2)
		kzFrameRight3 = images.Mirror(kzFrameLeft3)

		// Damaged
		kzFrameDamagedLeft1 = load("assets/sprites/KidZombieDamaged/left1.png")
		kzFrameDamagedLeft2 = load("assets/sprites/KidZombieDamaged/left2.png")
		kzFrameDamagedLeft3 = load("assets/sprites/KidZombieDamaged/left3.png")
		kzFrameDamagedUp1 = load("assets/sprites/KidZombieDamaged/up1.png")
		kzFrameDamagedUp2 = load("assets/sprites/KidZombieDamaged/up2.png")
		kzFrameDamagedUp3 = load("assets/sprites/KidZombieDamaged/up3.png")
		kzFrameDamagedDown1 = load("assets/sprites/KidZombieDamaged/down1.png")
		kzFrameDamagedDown2 = load("assets/sprites/KidZombieDamaged/down2.png")
		kzFrameDamagedDown3 = load("assets/sprites/KidZombieDamaged/down3.png")
		kzFrameDamagedRight1 = images.Mirror(kzFrameDamagedLeft1)
		kzFrameDamagedRight2 = images.Mirror(kzFrameDamagedLeft2)
		kzFrameDamagedRight3 = images.Mirror(kzFrameDamagedLeft3)

	})
}
func generateKidZombieVelicity() *components.Velocity {
	return &components.Velocity{
		Val: KZ_VELOCITY_LOWER_BOUND + rand.Float64()*KZ_VELOCITY_HIGHER_BOUND,
	}
}
