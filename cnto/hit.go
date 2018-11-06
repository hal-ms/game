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
		fmt.Println(p)
		point := repo.Hit.Add(p)
		// 叩いてる
		if p > 30 {
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
		// ステージ3なら
		if point < 600 {
			service.LCD.Next(2)
		} else if point < 400 {
			service.LCD.Next(1)
		} else if point < 200 {
			service.LCD.Next(0)
		}
	} else {
		service.LCD.Hide()
	}
}
