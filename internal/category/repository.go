package category

// Repository defines data-access methods for categories.
type Repository interface {
	FindAll() ([]Category, error)
	FindByID(id string) (*Category, error)
	Create(c *Category) error
}

// MemoryRepository stores categories in a map keyed by category ID.
type MemoryRepository struct {
	categories map[string]*Category
}

// NewMemoryRepository creates an empty in-memory category store.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		categories: make(map[string]*Category),
	}
}

func (r *MemoryRepository) FindAll() ([]Category, error) {
	// TODO: implement list all categories
	return nil, nil
}

func (r *MemoryRepository) FindByID(id string) (*Category, error) {
	// TODO: implement find by ID
	return nil, nil
}

func (r *MemoryRepository) Create(c *Category) error {
	// TODO: implement create
	return nil
}
