package wallet

// Transaction records a wallet credit or debit for a user.
// Type values will typically be: "credit", "debit".
type Transaction struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}
