package feedback

import(
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Feedback struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    bson.ObjectID `bson:"userId" json:"userId"`
	Rating    int           `bson:"rating" json:"rating"`
	Message   string        `bson:"message" json:"message"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

