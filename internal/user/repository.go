package user

import (
	"context"

	"github.com/amanhasnainy/bingo-backend/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository struct {
}

func (r *Repository) FindByEmail(
	email string,
) (*User, error) {

	collection := database.DB.Collection("users")

	var user User

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"email": email,
		},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Create(
	user User,
) error {

	collection := database.DB.Collection("users")

	_, err := collection.InsertOne(
		context.Background(),
		user,
	)

	return err
}

func (r *Repository) FindByID(
	id string,
) (*User, error) {

	collection := database.DB.Collection(
		"users",
	)

	objectID, err := bson.ObjectIDFromHex(
		id,
	)

	if err != nil {
		return nil, err
	}

	var user User

	err = collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(
		&user,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}