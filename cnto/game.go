package cnto

import (
	"github.com/hal-ms/game/log"
	"github.com/hal-ms/game/repo"
	"github.com/hal-ms/game/service"
)

func Game(p int) {
	if p > 10 {
		if !repo.State.Get().IsHit {
			err := service.LCD.Start()
			if err != nil {
				log.SendSlack(err.Error())
				return
			}
			repo.State.IsHit(true)
		}
	} else {
		if repo.State.Get().IsHit {
			err := service.LCD.Stop()
			if err != nil {
				log.SendSlack(err.Error())
				return
			}
			repo.State.IsHit(false)
		}
	}
}

func Stage(p int) {
	if p > 600 {
		service.LCD.Next(2)
	} else if p > 200 {
		service.LCD.Next(1)
	} else if p > 60 {
		service.LCD.Next(0)
	}
	
}
