package main

import (
	"fmt"
	"log"

	"github.com/emilieanthony/senzr/internal/db"
	"github.com/emilieanthony/senzr/internal/svc/rpi"
	"github.com/gin-gonic/gin"
)

const (
	Port     = 3000
	BasePath = "/api"
	V1       = "/v1"
)

func main() {
	if err := db.MigrateSchemas(); err != nil {
		log.Fatal(err.Error())
	}

	database := db.Db{}

	// controllers
	rpiServer := rpi.Server{Db: &database}

	// routes
	r := gin.Default()
	v1 := r.Group(BasePath + V1)
	{
		v1.GET("/helloworld", rpiServer.HelloWorld)
		v1.GET("/co2/latest", rpiServer.GetLatestCarbonDioxideEntry)
	}

	// start app
	err := r.Run(fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("server successfully started on port %v", Port)
}
