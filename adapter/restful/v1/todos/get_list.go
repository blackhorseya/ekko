package todos

import (
	"github.com/gin-gonic/gin"
)

// GetList is the api to get todo list.
// @Summary Get todo list.
// @Description Get todo list.
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]agg.Issue}
// @Failure 400,404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/todos [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/4/15|sean|implement me
	panic("implement me")
}
