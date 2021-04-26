package carbon_dioxide

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) HelloWorld(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world!")
}
