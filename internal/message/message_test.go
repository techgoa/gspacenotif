package message

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/techgoa/gspacenotif/internal/types"
)

func TestFormatProductErrorMessage(t *testing.T) {
	tests := []struct {
		name          string
		ecommerceName string
		params        types.ProductErrorParams
		wantContains  []string
	}{
		{
			name:          "Normal message",
			ecommerceName: "TestStore",
			params: types.ProductErrorParams{
				Title:             "Test Error",
				Error:             "Error message",
				ShopID:            "123",
				ProductMerchantID: "456",
				Response:          "Test response",
			},
			wantContains: []string{
				"Test Error",
				"TestStore",
				"123",
				"456",
				"Test response",
				"Error message",
			},
		},
		{
			name:          "Empty ecommerce name",
			ecommerceName: "",
			params: types.ProductErrorParams{
				Title: "Test",
				Error: "Error",
			},
			wantContains: []string{
				"Test",
				"Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatProductErrorMessage(tt.ecommerceName, tt.params)

			for _, want := range tt.wantContains {
				if !strings.Contains(result, want) {
					t.Errorf("FormatProductErrorMessage() = %v, should contain %v", result, want)
				}
			}
		})
	}
}

func TestCreateGoogleSpacesPayload(t *testing.T) {
	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{
			name:    "Valid message",
			message: "Test message",
			wantErr: false,
		},
		{
			name:    "Empty message",
			message: "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := CreateGoogleSpacesPayload(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGoogleSpacesPayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var result types.MessagePayload
			if err := json.Unmarshal(payload, &result); err != nil {
				t.Errorf("Failed to unmarshal payload: %v", err)
				return
			}

			if result.Text != tt.message {
				t.Errorf("Expected message %q, got %q", tt.message, result.Text)
			}
		})
	}
}
