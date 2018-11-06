package cnto

import (
	"fmt"

	"github.com/hal-ms/game/log"
	"github.com/hal-ms/game/repo"
	"github.com/hal-ms/game/service"
)

func Hit(p int) {
	if repo.State.Get().IsStandby {
		return
	}
	if repo.State.Get().IsWearing {
		point := repo.Hit.Add(p)
		// 叩いてる
		if p > 3 {
			err := service.LCD.Start()
			if err != nil {
				log.SendSlack(err.Error())
				return
			}
		} else {
			err := service.LCD.Stop()
			if err != nil {
				log.SendSlack(err.Error())
				return
			}
		}
		fmt.Println(point)
		// ステージ3なら
		if point < 60 {
			service.LCD.Next(2)
		} else if point < 40 {
			service.LCD.Next(1)
		} else if point < 20 {
			service.LCD.Next(0)
		}
	} else {
		service.LCD.Hide()
	}
}
