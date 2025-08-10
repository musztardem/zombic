package systems

import (
	"math"

	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/entities"
)

func FindNearestEnemyPosition(player *entities.Player, enemies *[]entities.EnemyBehaviour) *components.Position {
	var nearestEnemyPosition *components.Position
	var distance float64 = math.MaxFloat64

	for _, enemy := range *enemies {
		currentDistance := player.Position.DistanceTo(enemy.GetPosition())
		if currentDistance < distance {
			distance = currentDistance
			nearestEnemyPosition = enemy.GetPosition()
		}
	}

	return nearestEnemyPosition
}
