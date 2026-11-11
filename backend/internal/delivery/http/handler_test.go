package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"vespa-realtime-platform/backend/internal/domain"
	"vespa-realtime-platform/backend/internal/usecase"
)

func TestHandleUserEvent(t *testing.T) {
	uc := usecase.NewEventUsecase(nil, nil)
	h := NewHandler(uc)

	event := domain.UserEvent{
		EventID:   "evt-123",
		UserID:    "usr-456",
		Action:    "click",
		ProductID: "prod-789",
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/events", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	h.HandleUserEvent(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("expected status 202 Accepted, got %d", w.Code)
	}
}
