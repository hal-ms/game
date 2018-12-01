package cnto

import (
	"fmt"

	"github.com/hal-ms/game/repo"
	"github.com/hal-ms/game/service"
	"github.com/makki0205/log"
)

func Game(p int) {
	if p > 30 {
		if !repo.State.Get().IsHit {
			err := service.LCD.Start()
			if err != nil {
				log.Err(err)
				return
			}
			repo.State.IsHit(true)
			Stage(repo.Hit.Get().Point)
		}
	} else {
		if repo.State.Get().IsHit {
			err := service.LCD.Stop()
			if err != nil {
				log.Err(err)
				return
			}
			repo.State.IsHit(false)
		}
	}
}

func Stage(p int) {
	fmt.Println(p)
	if p > 9000 {
		repo.Ending = true
		err := service.LCD.Next(3)
		if err != nil {
			panic(err)
		}
		service.Main.End()
	} else if p > 6000 {
		err := service.LCD.Next(2)
		if err != nil {
			panic(err)
		}
	} else if p > 2500 {
		err := service.LCD.Next(1)
		if err != nil {
			panic(err)
		}
	}
}
