package gspacenotif

import "github.com/techgoa/gspacenotif/internal/http"

type NotifierConfig struct {
	WebhookURL      string
	EcommerceName   string
	LoggerFunc      func(level, source, payload, err string)
	LogLevelWarning string
}

type Notifier struct {
	config NotifierConfig
	client *http.Client
}
