package repository

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/fx"
)

// Rating enthält auch den Namen des Cocktails zum Zeitpunkt der Bewertung.
type Rating struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	CocktailID   uint          `bson:"cocktailId" json:"cocktailId"`
	CocktailName string        `bson:"cocktailName" json:"cocktailName"`
	Stars        int           `bson:"stars" json:"stars"`
	Comment      string        `bson:"comment" json:"comment"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
}

// RatingRepository speichert Bewertungen in einer eigenen MongoDB-Collection.
type RatingRepository struct {
	ratings *mongo.Collection
}

// NewRatingRepository verwendet standardmäßig die lokale MongoDB.
func NewRatingRepository(lc fx.Lifecycle) (*RatingRepository, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		// Docker kann als Replica-Set-Adresse einen nur intern erreichbaren Hostnamen melden.
		uri = "mongodb://localhost:27017/?directConnection=true"
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri).SetServerSelectionTimeout(5 * time.Second))
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{OnStop: client.Disconnect})
	return &RatingRepository{ratings: client.Database("cocktails_go").Collection("ratings")}, nil
}

// Insert vergibt ID und Zeitstempel auf dem Server.
func (r *RatingRepository) Insert(ctx context.Context, rating *Rating) error {
	rating.ID = bson.NewObjectID()
	rating.CreatedAt = time.Now().UTC().Truncate(time.Millisecond)
	_, err := r.ratings.InsertOne(ctx, rating)
	return err
}

// Recent lädt die letzten 20 Bewertungen. Die ID entscheidet bei gleichem Zeitstempel.
func (r *RatingRepository) Recent(ctx context.Context) ([]Rating, error) {
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}, {Key: "_id", Value: -1}}).SetLimit(20)
	cursor, err := r.ratings.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	ratings := make([]Rating, 0)
	err = cursor.All(ctx, &ratings)
	return ratings, err
}

// WatchInserts abonniert neue Dokumente, auch wenn sie außerhalb der App eingefügt werden.
func (r *RatingRepository) WatchInserts(ctx context.Context) (*mongo.ChangeStream, error) {
	filter := bson.D{{Key: "operationType", Value: "insert"}}
	return r.ratings.Watch(ctx, mongo.Pipeline{bson.D{{Key: "$match", Value: filter}}})
}
