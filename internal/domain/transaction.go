package domain

type transaction struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
