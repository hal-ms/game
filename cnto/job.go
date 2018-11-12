package cnto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/repo"
)

func Job(c *gin.Context) {
	job := c.Param("job")

	if repo.Job.Exist(job) {
		repo.Job.Job(job)
		c.JSON(http.StatusOK, job)
	} else {
		repo.Job.Job("")
		c.JSON(http.StatusBadRequest, "存在しない仕事です")
		return
	}

}
