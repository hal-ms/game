package cnto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/repo"
)

func Button(c *gin.Context) {
	if !repo.State.Get().IsStandby {
		c.JSON(http.StatusBadRequest, "ゲーム中です！")
		return
	}

	repo.State.IsStandby(false)
	c.JSON(http.StatusOK, "ok")
}
