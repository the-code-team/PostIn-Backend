package providers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

var (
	eventBusOnce sync.Once
	eventBusProp *EventBusProvider
)

type EventBusProvider struct {
	client *redis.Client
}

func GetEventBusProvider() *EventBusProvider {
	eventBusOnce.Do(func() {
		eventBusProp = &EventBusProvider{
			client: redis.NewClient(&redis.Options{
				Addr:     os.Getenv("REDIS_ADDR"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       1,
			}),
		}
	})

	return eventBusProp
}

func (p *EventBusProvider) Publish(channel string, message string) error {
	ctx := context.Background()
	err := p.client.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

func (p *EventBusProvider) Subscribe(channel string) *redis.PubSub {
	return p.client.Subscribe(context.Background(), channel)
}
