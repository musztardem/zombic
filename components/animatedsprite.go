package components

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimatedSprite struct {
	currentAnimationName string
	animations           map[string]*Animation
}

type Animation struct {
	frames            []*ebiten.Image
	speed             int
	frameTime         int
	currentFrameIndex int
}

func NewAnimatedSprite() *AnimatedSprite {
	return &AnimatedSprite{
		animations: make(map[string]*Animation, 0),
	}
}

func (as *AnimatedSprite) RegisterAnimation(name string, frames []*ebiten.Image, speed int) {
	as.animations[name] = &Animation{
		frames:            frames,
		speed:             speed,
		frameTime:         0,
		currentFrameIndex: 0,
	}
}

func (as *AnimatedSprite) Play(animationName string) {
	animation, ok := as.animations[animationName]
	if !ok {
		log.Fatalf("failed to find animation %v", animationName)
	}
	as.currentAnimationName = animationName

	animation.frameTime++
	if animation.frameTime > animation.speed {
		animation.currentFrameIndex = (animation.currentFrameIndex + 1) % len(animation.frames)
		animation.frameTime = 0
	}
}

func (as *AnimatedSprite) Get() *ebiten.Image {
	if as.currentAnimationName == "" {
		for k := range as.animations {
			as.currentAnimationName = k
			break
		}
	}

	currentAnimation := as.animations[as.currentAnimationName]
	return currentAnimation.frames[currentAnimation.currentFrameIndex]
}
