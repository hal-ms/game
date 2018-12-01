package cnto

import (
	"fmt"

	"github.com/hal-ms/game/repo"
	"github.com/hal-ms/game/service"
)

func Hit(p int) {
	if repo.State.Get().IsStandby {
		fmt.Println(p)
		if p > 55 && !repo.Ending {
			service.GameStart()
		}
		return
	}
	if repo.State.Get().IsWearing {
		service.LCD.Show()
		repo.Hit.Add(p)
		Game(p)
	} else {
		service.LCD.Hide()
	}
}
