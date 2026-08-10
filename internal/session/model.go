package session

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	UserID       bson.ObjectID `bson:"userId"`
	RefreshToken string        `bson:"refreshToken"`
	UserAgent    string        `bson:"userAgent"`
	IPAddress    string        `bson:"ipAddress"`
	ExpiresAt    time.Time     `bson:"expiresAt"`
	CreatedAt    time.Time     `bson:"createdAt"`
}