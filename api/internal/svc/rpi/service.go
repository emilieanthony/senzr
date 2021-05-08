package rpi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	SENSOR DATA INTERFACE:
	Co2         float64   `json:"co2"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
*/

type Controller struct{}

func (c *Controller) HelloWorld(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world!")
}
