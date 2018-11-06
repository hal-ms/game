package cnto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/repo"
)

func IsWearing(c *gin.Context) {

	if c.Param("IsWearing") != "true" && c.Param("IsWearing") != "false" {
		c.JSON(http.StatusBadRequest, "パラメータ不正")
		return
	}

	if c.Param("IsWearing") == "true" {
		repo.State.IsWearing(true)
		return
	} else if c.Param("IsWearing") == "false" {
		repo.State.IsWearing(false)
		return
	}
	c.JSON(http.StatusOK, "ok")

}
