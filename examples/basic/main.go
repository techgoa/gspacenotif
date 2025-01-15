// Package main demonstrates basic usage of the gspacenotif package
// for sending notifications to Google Chat Spaces.
package main

import (
	"log"

	"github.com/techgoa/gspacenotif"
)

// main demonstrates sending a product error notification using gspacenotif.
// It creates a new notifier with custom config and sends a product error message.
// The webhook URL should be replaced with your actual Google Spaces webhook URL.
func main() {
	// Initialize notifier with config
	notifier := gspacenotif.NewNotifier(gspacenotif.NotifierConfig{
		WebhookURL:      "https://chat.googleapis.com/v1/spaces/...", // Replace with your actual webhook URL
		EcommerceName:   "MyStore",                                   // Your ecommerce name
		LoggerFunc:      nil,                                         // Your logger function, defaults to DefaultLogger
		LogLevelWarning: "WARN",                                      // Your warning level
	})

	// Send product error notification
	err := notifier.SendProductError(gspacenotif.ProductErrorParams{
		Title:             "Product Update Failed",
		Error:             "Invalid price format",
		ShopID:            "12345",
		ProductMerchantID: "67890",
		Response:          "400 Bad Request",
	})

	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}
}
