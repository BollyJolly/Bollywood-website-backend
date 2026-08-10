package wallet

// Repository defines data-access methods for wallet transactions.
type Repository interface {
	Create(t *Transaction) error
	FindByUserID(userID string) ([]Transaction, error)
	GetBalance(userID string) (int, error)
}

// MemoryRepository stores transactions in a slice for now.
type MemoryRepository struct {
	transactions []Transaction
}

// NewMemoryRepository creates an empty in-memory wallet store.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		transactions: make([]Transaction, 0),
	}
}

func (r *MemoryRepository) Create(t *Transaction) error {
	// TODO: implement create transaction
	return nil
}

func (r *MemoryRepository) FindByUserID(userID string) ([]Transaction, error) {
	// TODO: implement find transactions by user
	return nil, nil
}

func (r *MemoryRepository) GetBalance(userID string) (int, error) {
	// TODO: implement balance calculation from transaction history
	return 0, nil
}
