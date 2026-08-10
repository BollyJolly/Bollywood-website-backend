package room

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Room struct {
	ID        bson.ObjectID   `bson:"_id,omitempty" json:"id"`
	Code      string          `bson:"code" json:"code"`
	PlaylistID string         `bson:"playlistId" json:"playlistId"`
	HostID    bson.ObjectID   `bson:"hostId" json:"hostId"`
	PlayerIDs []bson.ObjectID `bson:"playerIds" json:"playerIds"`

	CalledNumbers []int `bson:"calledNumbers" json:"calledNumbers"`

	Status string `bson:"status" json:"status"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}