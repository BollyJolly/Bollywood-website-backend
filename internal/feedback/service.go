package feedback

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(
	userID string,
	rating int,
	message string,
) error {

	userObjectID, err := bson.ObjectIDFromHex(
		userID,
	)

	if err != nil {
		return err
	}

	feedback := Feedback{
		UserID:    userObjectID,
		Rating:    rating,
		Message:   message,
		CreatedAt: time.Now(),
	}

	return s.repository.Create(
		feedback,
	)
}