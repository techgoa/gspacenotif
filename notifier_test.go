package gspacenotif

import (
	"testing"

	gspacenotif "github.com/techgoa/gspacenotif/types"
)

func TestNewNotifier(t *testing.T) {
	tests := []struct {
		name          string
		config        NotifierConfig
		wantEcommerce string
		wantLogLevel  string
	}{
		{
			name:          "Default values",
			config:        NotifierConfig{WebhookURL: "https://example.com"},
			wantEcommerce: "Undefined",
			wantLogLevel:  "WARN",
		},
		{
			name: "Custom values",
			config: NotifierConfig{
				WebhookURL:      "https://example.com",
				EcommerceName:   "TestStore",
				LogLevelWarning: "WARN",
			},
			wantEcommerce: "TestStore",
			wantLogLevel:  "WARN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notifier := NewNotifier(tt.config)

			if notifier.config.EcommerceName != tt.wantEcommerce {
				t.Errorf("NewNotifier() EcommerceName = %v, want %v",
					notifier.config.EcommerceName, tt.wantEcommerce)
			}

			if notifier.config.LogLevelWarning != tt.wantLogLevel {
				t.Errorf("NewNotifier() LogLevelWarning = %v, want %v",
					notifier.config.LogLevelWarning, tt.wantLogLevel)
			}
		})
	}
}

func TestSendProductError(t *testing.T) {
	var logCalled bool
	testLogger := func(level, source, payload, err string) {
		logCalled = true
	}

	notifier := NewNotifier(NotifierConfig{
		WebhookURL:    "https://example.com",
		EcommerceName: "TestStore",
		LoggerFunc:    testLogger,
	})

	err := notifier.SendProductError(gspacenotif.ProductErrorParams{
		Title:             "Test Error",
		Error:             "Test error message",
		ShopID:            "123",
		ProductMerchantID: "456",
		Response:          "Test response",
	})

	if err == nil {
		t.Error("SendProductError() expected error for invalid webhook URL")
	}

	if !logCalled {
		t.Error("SendProductError() logger was not called")
	}
}
