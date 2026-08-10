package session

import (
	"context"

	"github.com/amanhasnainy/bingo-backend/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository struct {
}

// Create inserts a new session into MongoDB.
// Create inserts a new session into MongoDB and returns the saved session.
func (r *Repository) Create(
	session Session,
) (*Session, error) {

	collection := database.DB.Collection("sessions")

	result, err := collection.InsertOne(
		context.Background(),
		session,
	)

	if err != nil {
		return nil, err
	}

	session.ID = result.InsertedID.(bson.ObjectID)

	return &session, nil
}

// FindByRefreshToken finds a session using its refresh token.
func (r *Repository) FindByRefreshToken(
	refreshToken string,
) (*Session, error) {

	collection := database.DB.Collection("sessions")

	var session Session

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"refreshToken": refreshToken,
		},
	).Decode(&session)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

// Delete removes a session by its ID.
func (r *Repository) Delete(
	id bson.ObjectID,
) error {

	collection := database.DB.Collection("sessions")

	_, err := collection.DeleteOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
	)

	return err
}