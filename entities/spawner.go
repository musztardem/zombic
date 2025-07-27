package entities

import (
	"math/rand"

	"github.com/musztardem/zombic/components"
)

type Spawner struct {
	pathFollowLoop  *components.PathFollowLoop
	enemies         *[]EnemyBehaviour
	playerPosition  *components.Position
	secondsInterval int
	timeCounter     int
}

func NewSpawner(path *components.Path, velocity *components.Velocity, playerPosition *components.Position, enemies *[]EnemyBehaviour, secondsInterval int) *Spawner {
	return &Spawner{
		pathFollowLoop:  components.NewPathFollowLoop(path, velocity),
		playerPosition:  playerPosition,
		enemies:         enemies,
		secondsInterval: secondsInterval,
	}
}

func (s *Spawner) Update() {
	s.pathFollowLoop.Update()
	s.timeCounter++
	counterLimit := s.secondsInterval * 60

	if s.timeCounter > counterLimit {
		*s.enemies = append(*s.enemies, s.createNewEnemy())

		s.timeCounter = 0
	}
}

func (s *Spawner) GetPosition() *components.Position {
	return s.pathFollowLoop.GetPosition()
}

// 40% chance to spawn skinny zombie
// 40% chance to spawn kid zombie
// 20% chance to spawn big zombie
func (s *Spawner) createNewEnemy() EnemyBehaviour {
	chance := rand.Float64()
	if chance < 0.4 {
		return NewSkinnyZombie(&components.Position{
			X: s.GetPosition().X,
			Y: s.GetPosition().Y,
		}, s.playerPosition)
	}

	if chance >= 0.4 && chance < 0.8 {
		return NewKidZombie(&components.Position{
			X: s.GetPosition().X,
			Y: s.GetPosition().Y,
		}, s.playerPosition)
	}

	return NewBigZombie(&components.Position{
		X: s.GetPosition().X,
		Y: s.GetPosition().Y,
	}, s.playerPosition)
}
