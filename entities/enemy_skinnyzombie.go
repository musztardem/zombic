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

type SkinnyZombie struct {
	*Enemy
}

const (
	SZ_VELOCITY_LOWER_BOUND  = 0.25
	SZ_VELOCITY_HIGHER_BOUND = 0.75
)

var (
	skinnyZombieFramesLoaded sync.Once

	szFrameLeft1, szFrameLeft2, szFrameLeft3    *ebiten.Image
	szFrameUp1, szFrameUp2, szFrameUp3          *ebiten.Image
	szFrameDown1, szFrameDown2, szFrameDown3    *ebiten.Image
	szFrameRight1, szFrameRight2, szFrameRight3 *ebiten.Image

	szFrameDamagedLeft1, szFrameDamagedLeft2, szFrameDamagedLeft3    *ebiten.Image
	szFrameDamagedUp1, szFrameDamagedUp2, szFrameDamagedUp3          *ebiten.Image
	szFrameDamagedDown1, szFrameDamagedDown2, szFrameDamagedDown3    *ebiten.Image
	szFrameDamagedRight1, szFrameDamagedRight2, szFrameDamagedRight3 *ebiten.Image
)

func NewSkinnyZombie(position, targetPosition *components.Position) *SkinnyZombie {
	skinnyZombie := &SkinnyZombie{
		Enemy: &Enemy{
			Position:       position,
			TargetPosition: targetPosition,
			Velocity:       generateSkinnyZombieVelocity(),
		},
	}

	skinnyZombie.loadAnimations()
	skinnyZombie.updateCollider()

	return skinnyZombie
}

func (sz *SkinnyZombie) loadAnimations() {
	sz.loadEnemyFramesOnce()

	enemyAnimatedSprite := components.NewAnimatedSprite()
	enemyAnimatedSprite.RegisterAnimation(
		"walk_right",
		[]*ebiten.Image{
			szFrameRight1,
			szFrameRight2,
			szFrameRight1,
			szFrameRight3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_left",
		[]*ebiten.Image{
			szFrameLeft1,
			szFrameLeft2,
			szFrameLeft1,
			szFrameLeft3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_up",
		[]*ebiten.Image{
			szFrameUp1,
			szFrameUp2,
			szFrameUp1,
			szFrameUp3,
		},
		7,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"walk_down",
		[]*ebiten.Image{
			szFrameDown1,
			szFrameDown2,
			szFrameDown1,
			szFrameDown3,
		},
		7,
	)

	enemyAnimatedSprite.RegisterAnimation(
		"damaged_right",
		[]*ebiten.Image{
			szFrameDamagedRight1,
			szFrameDamagedRight2,
			szFrameDamagedRight3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_left",
		[]*ebiten.Image{
			szFrameDamagedLeft1,
			szFrameDamagedLeft2,
			szFrameDamagedLeft3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_up",
		[]*ebiten.Image{
			szFrameDamagedUp1,
			szFrameDamagedUp2,
			szFrameDamagedUp3,
		},
		3,
	)
	enemyAnimatedSprite.RegisterAnimation(
		"damaged_down",
		[]*ebiten.Image{
			szFrameDamagedDown1,
			szFrameDamagedDown2,
			szFrameDamagedDown3,
		},
		3,
	)

	enemyAnimatedSprite.RegisterAnimation(
		"idle",
		[]*ebiten.Image{
			szFrameDown1,
		},
		20,
	)

	sz.AnimatedSprite = enemyAnimatedSprite
}

func (sz *SkinnyZombie) loadEnemyFramesOnce() {
	skinnyZombieFramesLoaded.Do(func() {
		load := func(path string) *ebiten.Image {
			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				log.Fatalf("failed to load image %v", err)
			}
			return img
		}

		// Walking
		szFrameLeft1 = load("assets/sprites/SkinnyZombie/left1.png")
		szFrameLeft2 = load("assets/sprites/SkinnyZombie/left2.png")
		szFrameLeft3 = load("assets/sprites/SkinnyZombie/left3.png")
		szFrameUp1 = load("assets/sprites/SkinnyZombie/up1.png")
		szFrameUp2 = load("assets/sprites/SkinnyZombie/up2.png")
		szFrameUp3 = load("assets/sprites/SkinnyZombie/up3.png")
		szFrameDown1 = load("assets/sprites/SkinnyZombie/down1.png")
		szFrameDown2 = load("assets/sprites/SkinnyZombie/down2.png")
		szFrameDown3 = load("assets/sprites/SkinnyZombie/down3.png")
		szFrameRight1 = images.Mirror(szFrameLeft1)
		szFrameRight2 = images.Mirror(szFrameLeft2)
		szFrameRight3 = images.Mirror(szFrameLeft3)

		// Damaged
		szFrameDamagedLeft1 = load("assets/sprites/SkinnyZombieDamaged/left1.png")
		szFrameDamagedLeft2 = load("assets/sprites/SkinnyZombieDamaged/left2.png")
		szFrameDamagedLeft3 = load("assets/sprites/SkinnyZombieDamaged/left3.png")
		szFrameDamagedUp1 = load("assets/sprites/SkinnyZombieDamaged/up1.png")
		szFrameDamagedUp2 = load("assets/sprites/SkinnyZombieDamaged/up2.png")
		szFrameDamagedUp3 = load("assets/sprites/SkinnyZombieDamaged/up3.png")
		szFrameDamagedDown1 = load("assets/sprites/SkinnyZombieDamaged/down1.png")
		szFrameDamagedDown2 = load("assets/sprites/SkinnyZombieDamaged/down2.png")
		szFrameDamagedDown3 = load("assets/sprites/SkinnyZombieDamaged/down3.png")
		szFrameDamagedRight1 = images.Mirror(szFrameDamagedLeft1)
		szFrameDamagedRight2 = images.Mirror(szFrameDamagedLeft2)
		szFrameDamagedRight3 = images.Mirror(szFrameDamagedLeft3)
	})
}

func generateSkinnyZombieVelocity() *components.Velocity {
	return &components.Velocity{
		Val: SZ_VELOCITY_LOWER_BOUND + rand.Float64()*SZ_VELOCITY_HIGHER_BOUND,
	}
}

// var _ EnemyBehaviour = (*SkinnyZombie)(nil)
