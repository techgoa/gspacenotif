// Package gspacenotif provides functionality for sending notifications to Google Chat Spaces.
// It supports sending formatted error messages with product details and custom configurations.
package gspacenotif

import (
	"encoding/json"
	"fmt"
)

// ProductErrorParams defines the structure for product error notification parameters.
type ProductErrorParams struct {
	Title             string
	Error             string
	ShopID            string
	ProductMerchantID string
	Response          string
}

// MessagePayload represents the structure of a Google Chat message.
type MessagePayload struct {
	Text string `json:"text"`
}

// FormatProductErrorMessage creates a formatted error message with product details.
func FormatProductErrorMessage(ecommerceName string, params ProductErrorParams) string {
	return fmt.Sprintf(`❌ *%s* ❌

*%s Merchant ID:*
%s

*Product Merchant ID:*
%s

*Response:*
%s

*Error:*
%s`,
		params.Title,
		ecommerceName,
		params.ShopID,
		params.ProductMerchantID,
		params.Response,
		params.Error,
	)
}

func CreateGoogleSpacesPayload(message string) ([]byte, error) {
	payload := map[string]interface{}{
		"text": message,
	}

	return json.Marshal(payload)
}
