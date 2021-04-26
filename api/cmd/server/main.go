package main

import (
	"fmt"
	"log"

	"github.com/emilieanthony/senzr/internal/svc/carbon_dioxide"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

const (
	PORT = 3000
)

type Environment struct {
	DBHost string `envconfig:"db_host"` // SENZR_DB_HOST
}

func main() {
	// load env
	loadEnv()

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

func loadEnv() *Environment {
	var e Environment
	err := envconfig.Process("senzr", &e)
	if err != nil {
		log.Fatalf("could not load environment config: %s", err.Error())
	}
	log.Print("environment variables successfully loaded")
	return &e
}
