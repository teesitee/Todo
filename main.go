package main

import (
	"context"
	"todo-with-gig/api"
	"todo-with-gig/repository"
	"todo-with-gig/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// go entry to start program
	// initial everything & run web service
	// famous lib to run web service in Golang is Gin

	opt := options.Client().ApplyURI("mongodb+srv://mydev:1234@cluster0.wwlsg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	TodoDB := client.Database("toto-with-gig").Collection("todo")
	repoInstance := repository.MongoRepository{
		Collection: TodoDB,
	}
	h := api.Handler{
		Repo: repoInstance,
	}

	r := router.NewRouter(h)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
