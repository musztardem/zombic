package systems

import (
	"fmt"

	"github.com/musztardem/zombic/entities"
)

// System find missles that haven't hit any target and removes them from the slice so they can be garbage collected
func MissleMissed(missles *[]entities.Missle) {
	for i := len(*missles) - 1; i >= 0; i-- {
		if (*missles)[i].TicksLived > (*missles)[i].Lifetime {
			fmt.Println("removing a missle!")
			(*missles)[i].IsRemovable = true
		} else {
			(*missles)[i].TicksLived++
		}
	}

	for i := len(*missles) - 1; i >= 0; i-- {
		if (*missles)[i].IsRemovable {
			*missles = append((*missles)[:i], (*missles)[i+1:]...)
		}
	}
}
