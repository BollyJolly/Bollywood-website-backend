package wallet

// Service defines business operations for wallets.
type Service interface {
	GetBalance(userID string) (int, error)
	GetHistory(userID string) ([]Transaction, error)
	Credit(userID string, amount int) (*Transaction, error)
	Debit(userID string, amount int) (*Transaction, error)
}

// WalletService contains wallet business logic.
type WalletService struct {
	repo Repository
}

// NewService wires a Repository into the wallet service.
func NewService(repo Repository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) GetBalance(userID string) (int, error) {
	// TODO: implement get balance
	return 0, nil
}

func (s *WalletService) GetHistory(userID string) ([]Transaction, error) {
	// TODO: implement get transaction history
	return nil, nil
}

func (s *WalletService) Credit(userID string, amount int) (*Transaction, error) {
	// TODO: implement credit (e.g. prize payout)
	return nil, nil
}

func (s *WalletService) Debit(userID string, amount int) (*Transaction, error) {
	// TODO: implement debit (e.g. room entry fee)
	return nil, nil
}
