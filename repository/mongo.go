package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func (r MongoRepository) InsertTodo(text, personName string) error {
	_, err := r.Collection.InsertOne(context.TODO(), Todo{
		Text:       text,
		PersonName: personName,
	})
	return err
}

func (r MongoRepository) GetAllTodos() ([]Todo, error) {
	cur, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var todos = []Todo{}
	if err := cur.All(context.TODO(), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r MongoRepository) GetTodos(Name string) ([]Todo, error) {
	cur, err := r.Collection.Find(context.TODO(), bson.M{"person_name": Name})
	if err != nil {
		return nil, err
	}
	var todos = []Todo{}
	if err := cur.All(context.TODO(), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r MongoRepository) PatTodos(Name, text string) error {
	_, err := r.Collection.UpdateOne(context.TODO(),
		bson.M{"person_name": Name}, bson.M{"$set": bson.M{"data": text}})
	if err != nil {
		return err
	}

	return err
}

func (r MongoRepository) DelTodos(Name string) error {
	_, err := r.Collection.DeleteMany(context.TODO(), bson.M{"person_name": Name})
	if err != nil {
		return err
	}

	return err
}

type Todo struct {
	Text       string `bson:"data"`
	PersonName string `bson:"person_name"`
}
