package category

// Service defines business operations for categories.
type Service interface {
	List() ([]Category, error)
	GetByID(id string) (*Category, error)
}

// CategoryService contains category business logic.
type CategoryService struct {
	repo Repository
}

// NewService wires a Repository into the category service.
func NewService(repo Repository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) List() ([]Category, error) {
	// TODO: implement list categories
	return nil, nil
}

func (s *CategoryService) GetByID(id string) (*Category, error) {
	// TODO: implement get category by ID
	return nil, nil
}
