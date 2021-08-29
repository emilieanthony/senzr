package cloud_receiver

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

const (
	CollectionCO2         = "carbon_dioxide"
	CollectionTemperature = "temperature"
	CollectionHumidity    = "humidity"
	credentialsPath       = "../credentials/senzr-313218-507b7a0a8637.json"
)

type CollectionData struct {
	Value     float64   `json:"value" firestore:"value"`
	Timestamp time.Time `json:"createdAt" firestore:"created_at"`
}

type Data struct {
	Co2         float64 `json:"co2"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

// Message can contain any of these fields:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type Message struct {
	Data []byte `json:"data"`
}

func SensorDataReceiver(ctx context.Context, m Message) (err error) {
	var data Data
	if err := json.Unmarshal(m.Data, &data); err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}
	client, err := NewClient(ctx)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer func() {
		client.Close()
	}()
	ts, err := time.Parse(time.RFC3339, data.Timestamp)
	if err != nil {
		return fmt.Errorf("parsing timestamp: %w", err)
	}
	co2 := CollectionData{
		Value:     data.Co2,
		Timestamp: ts,
	}
	temperature := CollectionData{
		Value:     data.Temperature,
		Timestamp: ts,
	}
	humidity := CollectionData{
		Value:     data.Humidity,
		Timestamp: ts,
	}
	if _, err := client.Collection(CollectionCO2).NewDoc().Create(ctx, co2); err != nil {
		return fmt.Errorf("writing co2 data: %w", err)
	}
	if _, err := client.Collection(CollectionTemperature).NewDoc().Create(ctx, temperature); err != nil {
		return fmt.Errorf("writing co2 data: %w", err)
	}
	if _, err := client.Collection(CollectionHumidity).NewDoc().Create(ctx, humidity); err != nil {
		return fmt.Errorf("writing co2 data: %w", err)
	}
	return nil
}

func NewClient(ctx context.Context) (*firestore.Client, error) {
	env, err := Env()
	if err != nil {
		return nil, fmt.Errorf("getting environment: %w", err)
	}
	client, err := firestore.NewClient(ctx, env.Project, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client, nil
}

type environment struct {
	Project string `envconfig:"project" default:"senzr-313218"` // SENZR_PROJECT
}

func Env() (*environment, error) {
	var e environment
	err := envconfig.Process("senzr", &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
