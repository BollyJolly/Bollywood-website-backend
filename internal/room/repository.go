package room

import (
	"context"

	"github.com/amanhasnainy/bingo-backend/internal/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repository struct {
}

func (r *Repository) Create(
	room Room,
) error {

	collection := database.DB.Collection(
		"rooms",
	)

	_, err := collection.InsertOne(
		context.Background(),
		room,
	)

	return err
}

func (r *Repository) FindByCode(
	code string,
) (*Room, error) {

	collection := database.DB.Collection(
		"rooms",
	)

	var room Room

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"code": code,
		},
	).Decode(
		&room,
	)

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *Repository) Update(
	room Room,
) error {

	collection := database.DB.Collection(
		"rooms",
	)

	_, err := collection.ReplaceOne(
		context.Background(),
		bson.M{
			"_id": room.ID,
		},
		room,
	)

	return err
}

func (r *Repository) FindAll() ([]Room, error) {

	collection := database.DB.Collection("rooms")

	cursor, err := collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	rooms := []Room{}

	err = cursor.All(
		context.Background(),
		&rooms,
	)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}