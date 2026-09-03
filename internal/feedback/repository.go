package feedback

import (
	"context"

	"github.com/amanhasnainy/bingo-backend/internal/database"
)

type Repository struct {
}

func (r *Repository) Create(
	feedback Feedback,
) error {

	collection := database.DB.Collection(
		"feedback",
	)

	_, err := collection.InsertOne(
		context.Background(),
		feedback,
	)

	return err
}