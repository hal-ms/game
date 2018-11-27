package cnto

import (
	"fmt"

	"github.com/hal-ms/game/log"
	"github.com/hal-ms/game/repo"
	"github.com/hal-ms/game/service"
)

func Game(p int) {
	if p > 30 {
		if !repo.State.Get().IsHit {
			err := service.LCD.Start()
			if err != nil {
				log.SendSlack(err.Error())
				return
			}
			repo.State.IsHit(true)
			Stage(repo.Hit.Get().Point)
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
	fmt.Println(p)
	if p > 15000 {
		err := service.LCD.Next(3)
		if err != nil {
			panic(err)
		}
		service.Main.End()
	} else if p > 9000 {
		err := service.LCD.Next(2)
		if err != nil {
			panic(err)
		}
	} else if p > 3000 {
		err := service.LCD.Next(1)
		if err != nil {
			panic(err)
		}
	}
}
