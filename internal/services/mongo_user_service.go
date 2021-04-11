package services

import (
	"context"
	"fmt"
	"reflect"
	"strings"

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
	query := bson.M{}
	cursor, er1 := p.Collection.Find(ctx, query)
	if er1 != nil {
		return nil, er1
	}
	var result []User
	er2 := cursor.All(ctx, &result)
	if er2 != nil {
		return nil, er2
	}
	return &result, nil
}

func (p *MongoUserService) Load(ctx context.Context, id string) (*User, error) {
	query := bson.M{"_id": id}
	result := p.Collection.FindOne(ctx, query)
	if result.Err() != nil {
		if strings.Compare(fmt.Sprint(result.Err()), "mongo: no documents in result") == 0 {
			return nil, nil
		} else {
			return nil, result.Err()
		}
	}
	user := User{}
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *MongoUserService) Insert(ctx context.Context, user *User) (int64, error) {
	_, err := p.Collection.InsertOne(ctx, user)
	if err != nil {
		errMsg := err.Error()
		if strings.Index(errMsg, "duplicate key error collection:") >= 0 {
			if strings.Index(errMsg, "dup key: { _id: ") >= 0 {
				return 0, nil
			} else {
				return -1, nil
			}
		} else {
			return 0, err
		}
	}
	return 1, nil
}

func (p *MongoUserService) Update(ctx context.Context, user *User) (int64, error) {
	query := bson.M{"_id": user.Id}
	updateQuery := bson.M{
		"$set": user,
	}
	result, err := p.Collection.UpdateOne(ctx, query, updateQuery)
	if result.ModifiedCount > 0 {
		return result.ModifiedCount, err
	} else if result.UpsertedCount > 0 {
		return result.UpsertedCount, err
	} else {
		return result.MatchedCount, err
	}
}

func (p *MongoUserService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	maps := m.MakeMapBson(userType)
	filter := m.BuildQueryByIdFromMap(user, "id")
	bson := m.MapToBson(user, maps)
	return m.PatchOne(ctx, p.Collection, bson, filter)
}

func (p *MongoUserService) Delete(ctx context.Context, id string) (int64, error) {
	query := bson.M{"_id": id}
	result, err := p.Collection.DeleteOne(ctx, query)
	if result == nil || err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}
