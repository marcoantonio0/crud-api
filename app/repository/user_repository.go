package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Birthday primitive.DateTime `json:"birthday,omitempty" validate:"required"`
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection}
}

func (r *UserRepository) CreateUser(user User) (interface{}, error) {

	doc, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}
	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) List() ([]User, error) {
	filter := bson.M{}
	result, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer result.Close(context.Background())

	var users []User
	for result.Next(context.Background()) {
		var user User
		if err := result.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserRepository) FindById(id string) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}

	options := options.FindOne()

	var user User
	err = s.collection.FindOne(context.Background(), filter, options).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
