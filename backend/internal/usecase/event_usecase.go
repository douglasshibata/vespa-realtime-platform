package usecase

import (
	"context"

	"vespa-realtime-platform/backend/internal/domain"
)

type EventPublisher interface {
	PublishEvent(ctx context.Context, key string, event interface{}) error
}

type VespaFeeder interface {
	FeedDocument(ctx context.Context, schema string, docID string, doc interface{}) error
}

type EventUsecase struct {
	publisher EventPublisher
	vespa     VespaFeeder
}

func NewEventUsecase(publisher EventPublisher, vespa VespaFeeder) *EventUsecase {
	return &EventUsecase{
		publisher: publisher,
		vespa:     vespa,
	}
}

func (u *EventUsecase) ProcessUserEvent(ctx context.Context, event domain.UserEvent) error {
	if u.publisher != nil {
		if err := u.publisher.PublishEvent(ctx, event.UserID, event); err != nil {
			return err
		}
	}
	return nil
}

func (u *EventUsecase) SyncProductToVespa(ctx context.Context, product domain.Product) error {
	if u.vespa != nil {
		return u.vespa.FeedDocument(ctx, "product", product.ID, product)
	}
	return nil
}
