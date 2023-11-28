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

type CatFact struct {
	Fact   string `bson:"fact" json:"fact"`
	Length int    `bson:"length" json:"length"`
}

type DbStorage interface {
	GetAll() ([]*CatFact, error)
	Put(fact *CatFact) error
}

type MongoStorage struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoStorage() (*MongoStorage, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return &MongoStorage{
		client:     client,
		database:   "catfact",
		collection: "facts",
	}, nil
}

func (store *MongoStorage) GetAll() ([]*CatFact, error) {
	coll := store.client.Database(store.database).Collection(store.collection)

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	results := []*CatFact{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (store *MongoStorage) Put(fact *CatFact) error {
	coll := store.client.Database(store.database).Collection(store.collection)
	_, err := coll.InsertOne(context.TODO(), fact)
	return err
}

type MongoServer struct {
	storage DbStorage
}

func NewMongoServer(d DbStorage) *MongoServer {
	return &MongoServer{
		storage: d,
	}
}

func (s *MongoServer) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	results, err := s.storage.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

type CatFactWorker struct {
	storage     DbStorage
	apiEndpoint string
}

func NewCatFactWorker(dbStorage DbStorage, apiEndpoint string) *CatFactWorker {
	return &CatFactWorker{
		storage:     dbStorage,
		apiEndpoint: apiEndpoint,
	}
}

func (cfw *CatFactWorker) start() error {
	ticker := time.NewTicker(2 * time.Second)

	for {
		resp, err := http.Get(cfw.apiEndpoint)
		if err != nil {
			return err
		}
		var catFact CatFact //map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		if err := cfw.storage.Put(&catFact); err != nil {
			return err
		}

		<-ticker.C
	}
}

func main() {
	mongoDbStore, err := NewMongoStorage()
	if err != nil {
		log.Fatal(err)
	}

	worker := NewCatFactWorker(mongoDbStore, "https://catfact.ninja/fact")
	go worker.start()

	server := NewMongoServer(mongoDbStore)
	http.HandleFunc("/facts", server.handleGetAllFacts)

	http.ListenAndServe(":3000", nil)
}
