package systems

import (
	"github.com/musztardem/zombic/entities"
)

func MissleHit(enemies *[]entities.EnemyBehaviour, missles *[]entities.Missle) {
	if len(*enemies) == 0 || len(*missles) == 0 {
		return
	}

	for i, missle := range *missles {
		for _, enemy := range *enemies {
			if enemy.GetCollider().CollidesWith(missle.Collider) {
				// enemy.MarkAsDead()
				enemy.MarkAsHit()
				(*missles)[i].IsRemovable = true
			}
		}
	}

	for i := len(*missles) - 1; i >= 0; i-- {
		if (*missles)[i].IsRemovable {
			*missles = append((*missles)[:i], (*missles)[i+1:]...)
		}
	}

	for i := len(*enemies) - 1; i >= 0; i-- {
		if (*enemies)[i].IsDead() {
			*enemies = append((*enemies)[:i], (*enemies)[i+1:]...)
		}
	}
}
