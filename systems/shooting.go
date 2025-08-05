package systems

import (
	"fmt"
	"math"

	"github.com/musztardem/zombic/components"
	"github.com/musztardem/zombic/entities"
)

// System responsible for finding an enemy that is in the shortest distance from the player
// and shooting at his direction
func ShootAtNearestEnemy(player *entities.Player, enemies *[]entities.EnemyBehaviour, missles *[]entities.Missle) {
	if len(*enemies) == 0 {
		return
	}

	if !player.Weapon.CanShoot() {
		return
	}

	var nearestEnemyPosition *components.Position
	var distance float64 = math.MaxFloat64

	for _, enemy := range *enemies {
		currentDistance := player.Position.DistanceTo(enemy.GetPosition())
		if currentDistance < distance {
			distance = currentDistance
			nearestEnemyPosition = enemy.GetPosition()
		}
	}

	newTargetPosition := &components.Position{
		X: nearestEnemyPosition.X,
		Y: nearestEnemyPosition.Y,
	}
	newMissle := player.Weapon.ShootAt(newTargetPosition)
	if newMissle != nil {
		*missles = append(*missles, *newMissle)
	} else {
		fmt.Println("unexpected shot happened")
	}
}
