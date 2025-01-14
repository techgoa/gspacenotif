package gspacenotif

type MessagePayload struct {
	Text string `json:"text"`
}

type ProductErrorParams struct {
	Title             string
	Error             string
	ShopID            string
	ProductMerchantID string
	Response          string
}
