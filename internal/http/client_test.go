package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SendMessage(t *testing.T) {
	tests := []struct {
		name       string
		serverResp int
		message    string
		wantErr    bool
	}{
		{
			name:       "Successful send",
			serverResp: http.StatusOK,
			message:    "Test message",
			wantErr:    false,
		},
		{
			name:       "Server error",
			serverResp: http.StatusInternalServerError,
			message:    "Test message",
			wantErr:    true,
		},
		{
			name:       "Bad request",
			serverResp: http.StatusBadRequest,
			message:    "Test message",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("Expected POST request, got %s", r.Method)
				}
				if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
					t.Errorf("Expected Content-Type: application/json, got %s", contentType)
				}

				w.WriteHeader(tt.serverResp)
			}))
			defer server.Close()

			client := NewClient()
			err := client.SendMessage(server.URL, tt.message)

			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
