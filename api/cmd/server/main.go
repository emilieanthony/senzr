package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/emilieanthony/senzr/internal/config"
	"github.com/emilieanthony/senzr/internal/svc/rpi"
	"github.com/gin-gonic/gin"
)

const (
	Port     = 3000
	BasePath = "/api"
	V1       = "/v1"
)

func main() {
	env, err := config.Env()
	if err != nil {
		log.Fatal(err.Error())
	}
	clerkClient, err := clerk.NewClient(env.AuthKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	// controllers
	rpiServer := rpi.Server{}

	// routes
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group(BasePath + V1)
	{
		v1.GET("/ping", rpiServer.Ping)
	}
	v1.Use(Auth(clerkClient))
	{
		v1.GET("/co2/latest", rpiServer.GetLatestCarbonDioxideEntry)
		v1.GET("/co2/duration", rpiServer.GetDurationAverageCarbonDioxide)
		v1.GET("/temperature/latest", rpiServer.GetLatestTemperatureEntry)
		v1.GET("/humidity/latest", rpiServer.GetLatestHumidityEntry)
	}

	// start app
	err = r.Run(fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("server successfully started on port %v", Port)
}

func Auth(client clerk.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// remove auth in debug mode
		if gin.Mode() == "debug" {
			c.Next()
			return
		}
		session, err := client.Verification().Verify(c.Request)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
		}
		c.Set("session", session)
		c.Next()
	}
}
