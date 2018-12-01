package service

import (
	"errors"
	"fmt"

	"github.com/hal-ms/game/repo"
	"github.com/makki0205/log"
)

func GameStart() {
	if !repo.State.Get().IsStandby {
		log.Err(errors.New("開始要求失敗しました"))
		return
	}
	fmt.Println(Main.Start())
	if Main.Start() {
		err := LCD.SetJob(repo.Job.Get().Job)
		if err != nil {
			panic(err)
		}
		repo.State.IsStandby(false)
	} else {
		repo.State.IsStandby(true)
		log.Err(errors.New("開始要求失敗しました"))
		return
	}
}
