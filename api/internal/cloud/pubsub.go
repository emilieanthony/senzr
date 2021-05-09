package cloud

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	topic     = "senzr_rpi_data"
	ProjectID = "senzr-313218"
)

type PubSubData struct {
	Co2         float64   `json:"co2"`
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	Timestamp   time.Time `json:"timestamp"`
}

type PubSub struct {
	client       *pubsub.Client
	subscription *pubsub.Subscription
	topic        *pubsub.Topic
}

func NewPubSubClient(ctx context.Context) (*PubSub, error) {
	client, err := pubsub.NewClient(ctx, ProjectID)
	if err != nil {
		return nil, fmt.Errorf("creating client: %w", err)
	}
	t := client.Topic(topic)

	sub, err := client.CreateSubscription(context.Background(), "sensor",
		pubsub.SubscriptionConfig{Topic: t})
	if err != nil {
		return nil, fmt.Errorf("creating pubsub client: %w", err)
	}

	return &PubSub{
		client:       client,
		subscription: sub,
		topic:        t,
	}, nil
}

func (p *PubSub) Receive(ctx context.Context, fn func(b []byte)) error {
	if err := p.subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fn(m.Data)
		m.Ack()
	}); err != nil {
		return fmt.Errorf("new message: %w", err)
	}
	return nil
}

func (p *PubSub) Stop() {
	p.topic.Stop()
}
