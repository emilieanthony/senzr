package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/emilieanthony/senzr/internal/cloud"
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

	ctx := context.Background()
	client, err := cloud.NewPubSubClient(ctx)
	if err != nil {
		log.Fatalf("pubsub: %v", err)
	}
	defer func() {
		client.Stop()
	}()
	if err := client.Receive(ctx, func(b []byte) {
		var data *cloud.PubSubData
		if err := json.Unmarshal(b, data); err != nil {
			fmt.Printf("could not unmarshal: %v", err)
		} else {
			fmt.Printf("received data: %v", data)
		}
	}); err != nil {
		log.Fatalf("could not start pubsub: %v", err)
	}
	// routes
	r := gin.Default()
	r.GET("/helloworld", cs.HelloWorld)

	// start app
	err = r.Run(fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("server successfully started on port %v", PORT)
}
