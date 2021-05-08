package main

import (
	"fmt"
	"log"

	"github.com/emilieanthony/senzr/internal/db"

	"github.com/emilieanthony/senzr/internal/svc/rpi"
	"github.com/gin-gonic/gin"
)

const (
	PORT = 3000
)

func main() {
	if err := db.MigrateSchemas(); err != nil {
		log.Fatal(err.Error())
	}

	// controllers
	cs := rpi.Controller{}

	// routes
	r := gin.Default()
	r.GET("/helloworld", cs.HelloWorld)

	// start app
	err := r.Run(fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("server successfully started on port %v", PORT)
}
