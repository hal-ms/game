package cnto

import (
	"github.com/hal-ms/game/repo"
	"github.com/hal-ms/game/service"
)

func Hit(p int) {
	if repo.State.Get().IsStandby {
		return
	}
	if repo.State.Get().IsWearing {
		service.LCD.Show()
		point := repo.Hit.Add(p)
		Game(p)
		Stage(point)
	} else {
		service.LCD.Hide()
	}
}
