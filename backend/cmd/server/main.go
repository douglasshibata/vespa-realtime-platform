package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	kafkaAdapter "vespa-realtime-platform/backend/internal/adapter/kafka"
	vespaAdapter "vespa-realtime-platform/backend/internal/adapter/vespa"
	handler "vespa-realtime-platform/backend/internal/delivery/http"
	"vespa-realtime-platform/backend/internal/domain"
	"vespa-realtime-platform/backend/internal/usecase"
)

func main() {
	kafkaBrokersEnv := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokersEnv == "" {
		kafkaBrokersEnv = "localhost:9092"
	}
	brokers := strings.Split(kafkaBrokersEnv, ",")

	vespaURLEnv := os.Getenv("VESPA_URL")
	if vespaURLEnv == "" {
		vespaURLEnv = "http://localhost:8080"
	}

	producer := kafkaAdapter.NewEventProducer(brokers, "user-events")
	defer producer.Close()

	vespaClient := vespaAdapter.NewVespaClient(vespaURLEnv)

	uc := usecase.NewEventUsecase(producer, vespaClient)
	h := handler.NewHandler(uc)

	consumer := kafkaAdapter.NewEventConsumer(brokers, "user-events", "vespa-indexer-group")
	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer.StartConsuming(ctx, func(ctx context.Context, event domain.UserEvent) error {
		log.Printf("[Event Indexer] Processing event %s for user %s", event.EventID, event.UserID)
		return nil
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/api/v1/events", h.HandleUserEvent)
	http.HandleFunc("/api/v1/products", h.HandleProductSync)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
