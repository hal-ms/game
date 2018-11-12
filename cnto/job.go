package cnto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-ms/game/repo"
)

func Job(c *gin.Context) {
	job := c.Param("job")

	j, _ := repo.Job.Exist(job)
	if j {
		repo.Job.Job(job)
		c.JSON(http.StatusOK, job)
	} else {
		repo.Job.Job("")
		c.JSON(http.StatusBadRequest, "存在しない仕事です")
		return
	}

}
