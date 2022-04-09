package service

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	. "go-service/internal/model"
)

type UserService interface {
	All(ctx context.Context) (*[]User, error)
	Load(ctx context.Context, id string) (*User, error)
	Insert(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
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
	filter := bson.M{}
	cursor, er1 := s.Collection.Find(ctx, filter)
	if er1 != nil {
		return nil, er1
	}
	var users []User
	er2 := cursor.All(ctx, &users)
	if er2 != nil {
		return nil, er2
	}
	return &users, nil
}

func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	filter := bson.M{"_id": id}
	res := s.Collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if strings.Compare(fmt.Sprint(res.Err()), "mongo: no documents in res") == 0 {
			return nil, nil
		} else {
			return nil, res.Err()
		}
	}
	user := User{}
	err := res.Decode(&user)
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
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": user,
	}
	res, err := s.Collection.UpdateOne(ctx, filter, update)
	if res.ModifiedCount > 0 {
		return res.ModifiedCount, err
	} else if res.UpsertedCount > 0 {
		return res.UpsertedCount, err
	} else {
		return res.MatchedCount, err
	}
}

func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	filter := bson.M{"_id": id}
	res, err := s.Collection.DeleteOne(ctx, filter)
	if res == nil || err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}
