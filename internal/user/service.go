package user

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(user User) error {

	user.Stars = 30

	return s.repository.Create(user)
}

func (s *Service) GetMe(
	userID string,
) (*User, error) {

	return s.repository.FindByID(
		userID,
	)
}

