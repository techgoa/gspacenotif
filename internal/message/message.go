package message

import (
	"encoding/json"
	"fmt"

	gspacenotif "github.com/techgoa/gspacenotif/types"
)

func FormatProductErrorMessage(ecommerceName string, params gspacenotif.ProductErrorParams) string {
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
