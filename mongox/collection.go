package mongox

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	c *mongo.Collection
}

func (p *Collection) Insert(document any) (string, error) {
	result, err := p.c.InsertOne(context.TODO(), document)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (p *Collection) InsertMany(documents ...any) ([]string, error) {
	result, err := p.c.InsertMany(context.TODO(), documents)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		ids[i] = id.(primitive.ObjectID).Hex()
	}
	return ids, nil
}

func (p *Collection) UpdateInsert(filter, update any) (string, error) {
	result, err := p.c.UpdateOne(context.TODO(), filter, update,
		options.Update().SetUpsert(true))
	if err != nil {
		return "", err
	}

	if result.UpsertedID != nil {
		return result.UpsertedID.(primitive.ObjectID).Hex(), nil
	}

	return "", nil
}

func (p *Collection) Update(id string, document any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.c.UpdateByID(context.Background(), objectID, bson.M{"$set": document})
	if err != nil {
		return err
	}

	return nil
}

func (p *Collection) UpdateMany(filter, update bson.M) error {
	_, err := p.c.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p *Collection) FindOne(id string, out any, fields ...string) error {
	var (
		opts          = options.FindOne()
		projection    = bson.M{}
		objectID, err = primitive.ObjectIDFromHex(id)
	)

	if err != nil {
		return err
	}

	for _, field := range fields {
		projection[field] = 1
	}

	result := p.c.FindOne(context.TODO(), bson.M{"_id": objectID}, opts)
	if result.Err() != nil {
		if !errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return result.Err()
		}
		return nil
	}

	return result.Decode(out)
}

// Find finds documents by filter and returns the results.
// The results must be a pointer to a slice.
func (p *Collection) Find(filter any, results any, fields ...string) error {
	var (
		opts       = options.Find()
		projection = bson.M{}
	)

	for _, field := range fields {
		projection[field] = 1
	}

	opts.SetProjection(projection)
	cursor, err := p.c.Find(context.TODO(), filter, opts)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		return err
	}

	return cursor.All(context.TODO(), results)
}

func (p *Collection) DeleteOne(filter any) (bool, error) {
	result, err := p.c.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return result.DeletedCount == 1, nil
}
