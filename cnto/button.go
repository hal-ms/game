package cnto

import (
	"fmt"
	"net/http"

	"github.com/hal-ms/game/service"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/repo"
)

func Button(c *gin.Context) {
	if !repo.State.Get().IsStandby {
		c.JSON(http.StatusBadRequest, "ゲーム中です！")
		return
	}
	fmt.Println(service.Main.Start())
	if service.Main.Start() {
		err := service.LCD.SetJob(repo.Job.Get().Job)
		if err != nil {
			panic(err)
		}
		repo.State.IsStandby(false)
		c.JSON(http.StatusOK, "ok")
	} else {
		repo.State.IsStandby(true)
		c.JSON(http.StatusBadRequest, "開始要求失敗しました")
		return
	}
}
