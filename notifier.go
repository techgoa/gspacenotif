package gspacenotif

import (
	"fmt"

	"github.com/techgoa/gspacenotif/internal/http"
	"github.com/techgoa/gspacenotif/internal/logger"
)

// NotifierConfig contains configuration options for the notifier.
type NotifierConfig struct {
	WebhookURL      string
	EcommerceName   string
	LoggerFunc      func(level, source, payload, err string)
	LogLevelWarning string
}

// Notifier handles sending notifications to Google Spaces.
type Notifier struct {
	config NotifierConfig
	client *http.Client
}

// NewNotifier creates a new Notifier instance with the provided configs.
// It sets default value for unspecified config options.
func NewNotifier(config NotifierConfig) *Notifier {
	if config.EcommerceName == "" {
		config.EcommerceName = "Undefined"
	}
	if config.LoggerFunc == nil {
		config.LoggerFunc = logger.DefaultLogger
	}
	if config.LogLevelWarning == "" {
		config.LogLevelWarning = "WARN"
	}

	return &Notifier{
		config: config,
		client: http.NewClient(),
	}
}

// SendProductError sends a formatted error message to Google Spaces
// with the provided product error parameters.
func (n *Notifier) SendProductError(params ProductErrorParams) error {
	formattedMessage := FormatProductErrorMessage(n.config.EcommerceName, params)

	googleSpacesRequest, err := CreateGoogleSpacesPayload(formattedMessage)
	if err != nil {
		n.config.LoggerFunc(n.config.LogLevelWarning,
			params.Title+" ERR PREPARING GOOGLE SPACES NOTIFICATION",
			err.Error(),
			err.Error())
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = n.client.SendMessage(n.config.WebhookURL, string(googleSpacesRequest))
	if err != nil {
		n.config.LoggerFunc(n.config.LogLevelWarning,
			params.Title+" ERR SEND TO GOOGLE SPACES",
			err.Error(),
			err.Error())
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
