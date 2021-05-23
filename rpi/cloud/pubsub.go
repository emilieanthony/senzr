package cloud

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/option"

	"cloud.google.com/go/pubsub"
	"github.com/emilieanthony/senzr/rpi/sensor/pico"
)

const (
	topic           = "senzr_rpi_data"
	credentialsPath = "../credentials/senzr-313218-1450f27a71a6.json"
	ProjectID       = "senzr-313218"
)

type PubSub struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func NewPubSubClient(ctx context.Context) (*PubSub, error) {
	client, err := pubsub.NewClient(ctx, ProjectID, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}
	topic := client.Topic(topic)
	return &PubSub{
		client: client,
		topic:  topic,
	}, nil
}

// Publish ignores errors and just log them instead
func (p *PubSub) Publish(ctx context.Context, data *pico.Data) {
	b, err := data.Encode()
	if err != nil {
		fmt.Printf("Error: parsing data into bytes: %v", err)
	} else {
		p.topic.Publish(ctx, &pubsub.Message{
			Data: b,
		})
		fmt.Printf("[%s] Published %d bytes to pubsub \n", time.Now().Format(time.RFC3339), len(b))
	}
}

func (p *PubSub) Stop() {
	p.topic.Stop()
}
