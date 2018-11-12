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
		service.LCD.Show()
		point := repo.Hit.Add(p)
		// 叩いてる
		if p > 3 {
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
		fmt.Println(point)
		// ステージ3なら
		if point > 600 {
			service.LCD.Next(2)
		} else if point > 200 {
			service.LCD.Next(1)
		} else if point > 60 {
			service.LCD.Next(0)
		}
	} else {
		service.LCD.Hide()
	}
}
