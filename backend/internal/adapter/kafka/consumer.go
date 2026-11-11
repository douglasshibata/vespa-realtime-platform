package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"vespa-realtime-platform/backend/internal/domain"
)

type EventConsumer struct {
	reader *kafka.Reader
}

func NewEventConsumer(brokers []string, topic string, groupID string) *EventConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10,
		MaxBytes: 10 * 1024 * 1024,
	})
	return &EventConsumer{reader: reader}
}

func (c *EventConsumer) StartConsuming(ctx context.Context, handler func(ctx context.Context, event domain.UserEvent) error) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				m, err := c.reader.ReadMessage(ctx)
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					log.Printf("[Kafka Consumer Error] %v", err)
					continue
				}

				var event domain.UserEvent
				if err := json.Unmarshal(m.Value, &event); err != nil {
					log.Printf("[Kafka Consumer] Failed to unmarshal message: %v", err)
					continue
				}

				if err := handler(ctx, event); err != nil {
					log.Printf("[Kafka Consumer] Failed to handle event: %v", err)
				}
			}
		}
	}()
}

func (c *EventConsumer) Close() error {
	return c.reader.Close()
}
