package types

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
