package db

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/emilieanthony/senzr/internal/config"
)

const (
	CollectionCO2         = "carbon_dioxide"
	CollectionTemperature = "temperature"
	CollectionHumidity    = "humidity"
	credentialsPath       = "../credentials/senzr-313218-507b7a0a8637.json"
)

func NewClient(ctx context.Context) (*firestore.Client, error) {
	env, err := config.Env()
	if err != nil {
		return nil, fmt.Errorf("getting environment: %w", err)
	}
	client, err := firestore.NewClient(ctx, env.Project, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client, nil
}
