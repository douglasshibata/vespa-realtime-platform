package http

import (
	"encoding/json"
	"net/http"
	"time"

	"vespa-realtime-platform/backend/internal/domain"
	"vespa-realtime-platform/backend/internal/usecase"
)

type Handler struct {
	usecase *usecase.EventUsecase
}

func NewHandler(usecase *usecase.EventUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) HandleUserEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event domain.UserEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if event.Timestamp == 0 {
		event.Timestamp = time.Now().Unix()
	}

	if err := h.usecase.ProcessUserEvent(r.Context(), event); err != nil {
		http.Error(w, "Failed to process event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"status": "queued", "eventId": event.EventID})
}

func (h *Handler) HandleProductSync(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid product payload", http.StatusBadRequest)
		return
	}

	if err := h.usecase.SyncProductToVespa(r.Context(), product); err != nil {
		http.Error(w, "Failed to sync product to Vespa", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "synced", "productId": product.ID})
}
