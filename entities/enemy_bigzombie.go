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

type BigZombie struct {
	*Enemy
}

const (
	BZ_VELOCITY_LOWER_BOUND  = 0.1
	BZ_VELOCITY_HIGHER_BOUND = 0.4
)

var (
	bigZombieFramesLoaded sync.Once

	bzFrameLeft1, bzFrameLeft2, bzFrameLeft3    *ebiten.Image
	bzFrameUp1, bzFrameUp2, bzFrameUp3          *ebiten.Image
	bzFrameDown1, bzFrameDown2, bzFrameDown3    *ebiten.Image
	bzFrameRight1, bzFrameRight2, bzFrameRight3 *ebiten.Image

	bzFrameDamagedLeft1, bzFrameDamagedLeft2, bzFrameDamagedLeft3    *ebiten.Image
	bzFrameDamagedUp1, bzFrameDamagedUp2, bzFrameDamagedUp3          *ebiten.Image
	bzFrameDamagedDown1, bzFrameDamagedDown2, bzFrameDamagedDown3    *ebiten.Image
	bzFrameDamagedRight1, bzFrameDamagedRight2, bzFrameDamagedRight3 *ebiten.Image
)

func NewBigZombie(position, targetPosition *components.Position) *BigZombie {
	bigZombie := &BigZombie{
		Enemy: &Enemy{
			Position:       position,
			TargetPosition: targetPosition,
			Velocity:       generateBigZombieVelocity(),
		},
	}

	bigZombie.loadAnimations()
	bigZombie.updateCollider()

	return bigZombie
}

func (bz *BigZombie) loadAnimations() {
	bz.loadEnemyFramesOnce()

	enemyAnimatedSprite := components.NewAnimatedSprite()
	enemyAnimatedSprite.RegisterAnimation(
		"walk_right",
		[]*ebiten.Image{
			bzFrameRight1,
			bzFrameRight2,
			bzFrameRight1,
			bzFrameRight3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_left",
		[]*ebiten.Image{
			bzFrameLeft1,
			bzFrameLeft2,
			bzFrameLeft1,
			bzFrameLeft3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_up",
		[]*ebiten.Image{
			bzFrameUp1,
			bzFrameUp2,
			bzFrameUp1,
			bzFrameUp3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_down",
		[]*ebiten.Image{
			bzFrameDown1,
			bzFrameDown2,
			bzFrameDown1,
			bzFrameDown3,
		},
		7,
	)

	enemyAnimatedSprite.RegisterAnimation(
		"damaged_right",
		[]*ebiten.Image{
			bzFrameDamagedRight1,
			bzFrameDamagedRight2,
			bzFrameDamagedRight3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_left",
		[]*ebiten.Image{
			bzFrameDamagedLeft1,
			bzFrameDamagedLeft2,
			bzFrameDamagedLeft3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_up",
		[]*ebiten.Image{
			bzFrameDamagedUp1,
			bzFrameDamagedUp2,
			bzFrameDamagedUp3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_down",
		[]*ebiten.Image{
			bzFrameDamagedDown1,
			bzFrameDamagedDown2,
			bzFrameDamagedDown3,
		},
		3,
	)

	enemyAnimatedSprite.RegisterAnimation(
		"idle",
		[]*ebiten.Image{
			bzFrameDown1,
		},
		20,
	)

	bz.AnimatedSprite = enemyAnimatedSprite
}

func (bz *BigZombie) loadEnemyFramesOnce() {
	bigZombieFramesLoaded.Do(func() {
		load := func(path string) *ebiten.Image {
			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				log.Fatalf("failed to load image %v", err)
			}
			return img
		}

		// Walking
		bzFrameLeft1 = load("assets/sprites/BigZombie/left1.png")
		bzFrameLeft2 = load("assets/sprites/BigZombie/left2.png")
		bzFrameLeft3 = load("assets/sprites/BigZombie/left3.png")
		bzFrameUp1 = load("assets/sprites/BigZombie/up1.png")
		bzFrameUp2 = load("assets/sprites/BigZombie/up2.png")
		bzFrameUp3 = load("assets/sprites/BigZombie/up3.png")
		bzFrameDown1 = load("assets/sprites/BigZombie/down1.png")
		bzFrameDown2 = load("assets/sprites/BigZombie/down2.png")
		bzFrameDown3 = load("assets/sprites/BigZombie/down3.png")
		bzFrameRight1 = images.Mirror(bzFrameLeft1)
		bzFrameRight2 = images.Mirror(bzFrameLeft2)
		bzFrameRight3 = images.Mirror(bzFrameLeft3)

		// Damaged
		bzFrameDamagedLeft1 = load("assets/sprites/BigZombieDamaged/left1.png")
		bzFrameDamagedLeft2 = load("assets/sprites/BigZombieDamaged/left2.png")
		bzFrameDamagedLeft3 = load("assets/sprites/BigZombieDamaged/left3.png")
		bzFrameDamagedUp1 = load("assets/sprites/BigZombieDamaged/up1.png")
		bzFrameDamagedUp2 = load("assets/sprites/BigZombieDamaged/up2.png")
		bzFrameDamagedUp3 = load("assets/sprites/BigZombieDamaged/up3.png")
		bzFrameDamagedDown1 = load("assets/sprites/BigZombieDamaged/down1.png")
		bzFrameDamagedDown2 = load("assets/sprites/BigZombieDamaged/down2.png")
		bzFrameDamagedDown3 = load("assets/sprites/BigZombieDamaged/down3.png")
		bzFrameDamagedRight1 = images.Mirror(bzFrameDamagedLeft1)
		bzFrameDamagedRight2 = images.Mirror(bzFrameDamagedLeft2)
		bzFrameDamagedRight3 = images.Mirror(bzFrameDamagedLeft3)
	})
}

func generateBigZombieVelocity() *components.Velocity {
	return &components.Velocity{
		Val: BZ_VELOCITY_LOWER_BOUND + rand.Float64()*BZ_VELOCITY_HIGHER_BOUND,
	}
}
