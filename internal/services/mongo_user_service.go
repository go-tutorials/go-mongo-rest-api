package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	m "github.com/common-go/mongo"

	. "go-service/internal/models"
)

type MongoUserService struct {
	Collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *MongoUserService {
	collectionName := "users"
	return &MongoUserService{Collection: db.Collection(collectionName)}
}

func (p *MongoUserService) GetAll(ctx context.Context) (*[]User, error) {
	var result []User
	_, err := m.FindAndDecode(ctx, p.Collection, bson.M{}, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *MongoUserService) Load(ctx context.Context, id string) (*User, error) {
	var user User
	query := bson.M{"_id": id}
	ok, err := m.FindOneAndDecode(ctx, p.Collection, query, &user)
	if ok {
		return &user, err
	}
	return nil, err
}

func (p *MongoUserService) Insert(ctx context.Context, user *User) (int64, error) {
	return m.InsertOne(ctx, p.Collection, user)
}

func (p *MongoUserService) Update(ctx context.Context, user *User) (int64, error) {
	query := bson.M{"_id": user.Id}
	return m.UpdateOne(ctx, p.Collection, user, query)
}

func (p *MongoUserService) Delete(ctx context.Context, id string) (int64, error) {
	query := bson.M{"_id": id}
	return m.DeleteOne(ctx, p.Collection, query)
}
