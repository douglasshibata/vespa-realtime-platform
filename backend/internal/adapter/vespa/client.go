package vespa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type VespaClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewVespaClient(baseURL string) *VespaClient {
	return &VespaClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (v *VespaClient) FeedDocument(ctx context.Context, schema string, docID string, doc interface{}) error {
	endpoint := fmt.Sprintf("%s/document/v1/%s/%s/docid/%s", v.baseURL, schema, schema, docID)

	payload, err := json.Marshal(map[string]interface{}{
		"fields": doc,
	})
	if err != nil {
		return fmt.Errorf("marshal doc error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("create request error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("vespa document API error: status %d", resp.StatusCode)
	}

	return nil
}
