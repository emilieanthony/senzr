package main

import (
	"fmt"
	"log"

	"github.com/emilieanthony/senzr/internal/svc/carbon_dioxide"
	"github.com/gin-gonic/gin"
)

const (
	PORT = 3000
)

func main() {
	// TODO: add database migrations
	// https://github.com/golang-migrate/migrate

	// controllers
	cs := carbon_dioxide.Controller{}

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
