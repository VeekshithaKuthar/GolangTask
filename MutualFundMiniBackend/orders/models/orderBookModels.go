package models

type OrderBookResponse struct {
	ID          uint    `json:"id"`
	SchemeCode  string  `json:"scheme_code"`
	Side        string  `json:"side"`
	Amount      float64 `json:"amount"`
	Units       float64 `json:"units"`
	Status      string  `json:"status"`
	ContractURL string  `json:"contract_url,omitempty"`
}
