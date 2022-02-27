package service

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	mgo "github.com/core-go/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	. "go-service/internal/model"
)

type UserService interface {
	All(ctx context.Context) (*[]User, error)
	Load(ctx context.Context, id string) (*User, error)
	Insert(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type userService struct {
	Collection *mongo.Collection
}

func NewUserService(db *mongo.Database) UserService {
	collectionName := "users"
	return &userService{Collection: db.Collection(collectionName)}
}

func (s *userService) All(ctx context.Context) (*[]User, error) {
	query := bson.M{}
	cursor, er1 := s.Collection.Find(ctx, query)
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

func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	query := bson.M{"_id": id}
	result := s.Collection.FindOne(ctx, query)
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

func (s *userService) Insert(ctx context.Context, user *User) (int64, error) {
	_, err := s.Collection.InsertOne(ctx, user)
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

func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	query := bson.M{"_id": user.Id}
	updateQuery := bson.M{
		"$set": user,
	}
	result, err := s.Collection.UpdateOne(ctx, query, updateQuery)
	if result.ModifiedCount > 0 {
		return result.ModifiedCount, err
	} else if result.UpsertedCount > 0 {
		return result.UpsertedCount, err
	} else {
		return result.MatchedCount, err
	}
}

func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	maps := mgo.MakeBsonMap(userType)
	filter := mgo.BuildQueryByIdFromMap(user, "id")
	bson := mgo.MapToBson(user, maps)
	return mgo.PatchOne(ctx, s.Collection, bson, filter)
}

func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	query := bson.M{"_id": id}
	result, err := s.Collection.DeleteOne(ctx, query)
	if result == nil || err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}
