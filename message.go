package gspacenotif

import (
	"encoding/json"
	"fmt"
)

type ProductErrorParams struct {
	Title             string
	Error             string
	ShopID            string
	ProductMerchantID string
	Response          string
}

type MessagePayload struct {
	Text string `json:"text"`
}

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
