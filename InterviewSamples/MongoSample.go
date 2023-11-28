package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type MongoServer struct {
	client *mongo.Client
}

func NewMongoServer(c *mongo.Client) *MongoServer {
	return &MongoServer{
		client: c,
	}
}

func (s *MongoServer) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	coll := s.client.Database("catfact").Collection("facts")

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

type CatFactWorker struct {
	client *mongo.Client
}

func NewCatFactWorker(c *mongo.Client) *CatFactWorker {
	return &CatFactWorker{
		client: c,
	}
}

func (cfw *CatFactWorker) start() error {
	coll := cfw.client.Database("catfact").Collection("facts")
	ticker := time.NewTicker(2 * time.Second)

	for {
		resp, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}
		var catFact bson.M //map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		_, err = coll.InsertOne(context.TODO(), catFact)
		if err != nil {
			return err
		}

		<-ticker.C
	}
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	worker := NewCatFactWorker(client)
	go worker.start()

	server := NewMongoServer(client)
	http.HandleFunc("/facts", server.handleGetAllFacts)

	http.ListenAndServe(":3000", nil)
}
