package mongodb

import (
	"cinema.cassia.io/users/pkg/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	C *mongo.Collection
}

func (m *UserModel) All() ([]models.User, error) {
	ctx := context.TODO()
	var uu []models.User

	// Find all users
	userCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = userCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}
	return uu, err
}

func (m *UserModel) FindByID(id string) (*models.User, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user = models.User{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) Insert(user models.User) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}

func (m *UserModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
