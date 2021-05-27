package mongodb

import (
	"cinema.cassia.io/showtimes/pkg/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShowTimeModel struct {
	C *mongo.Collection
}

func (m *ShowTimeModel) All() ([]models.ShowTime, error) {
	ctx := context.TODO()
	var st []models.ShowTime

	showTimeCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = showTimeCursor.All(ctx, &st)
	if err != nil {
		return nil, err
	}
	return st, err
}

func (m *ShowTimeModel) FindByID(id string) (*models.ShowTime, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var showtime = models.ShowTime{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&showtime)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &showtime, nil
}

func (m *ShowTimeModel) FindByDate(date string) (*models.ShowTime, error) {
	var showtime = models.ShowTime{}
	err := m.C.FindOne(context.TODO(), bson.M{"date": date}).Decode(&showtime)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}
	return &showtime, nil
}

func (m *ShowTimeModel) Insert(showTime models.ShowTime) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), showTime)
}

func (m *ShowTimeModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}