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

	repo.State.IsStandby(false)
	service.LCD.SetJob()
	c.JSON(http.StatusOK, "ok")
}
