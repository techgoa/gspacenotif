package gspacenotif

import (
	"encoding/json"
	"testing"
)

func TestMessagePayloadJSON(t *testing.T) {
	payload := MessagePayload{
		Text: "Test message",
	}

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal MessagePayload: %v", err)
	}

	var unmarshaled MessagePayload
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal MessagePayload: %v", err)
	}

	if unmarshaled.Text != payload.Text {
		t.Errorf("Expected text %q, got %q", payload.Text, unmarshaled.Text)
	}
}
