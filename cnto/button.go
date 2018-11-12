package cnto

import (
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

	if repo.Job.Get().Job == "" {
		c.JSON(http.StatusBadRequest, "仕事がありません！")
		return
	}

	repo.State.IsStandby(false)

	if service.Main.Start() {
		service.LCD.SetJob(repo.Job.Get().Job)
		c.JSON(http.StatusOK, "ok")
	} else {
		repo.State.IsStandby(true)
		c.JSON(http.StatusBadRequest, "開始要求失敗しました")
		return
	}
}
