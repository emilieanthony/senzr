package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/emilieanthony/senzr/internal/config"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	CollectionCO2         = "carbon_dioxide"
	CollectionTemperature = "temperature"
	CollectionHumidity    = "humidity"
)

func NewClient(ctx context.Context) (*firestore.Client, error) {
	env, err := config.Env()
	if err != nil {
		return nil, fmt.Errorf("getting environment: %w", err)
	}
	client, err := firestore.NewClient(ctx, env.Project)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client, nil
}
