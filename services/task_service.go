package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"../tasks"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb://localhost:27017"

const dbName = "test"

const collectionName = "newCollection"

type testTask struct {
	_id       primitive.ObjectID
	item      string
	completed bool
}

// collection object/instance
var collection *mongo.Collection

//create connection with mongo db
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// check connection:
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Db initialization succeeded, collection created")

}

// GetTaskList : get all tasks in db
func GetTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cursor.Next(context.Background()) {
		var result bson.M
		e := cursor.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println("cursor..>", cursor, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)
	}

	fmt.Println("len(results): %d", len(results))

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.Background())
	json.NewEncoder(w).Encode(results)
}

// GetTask : for GET http method, expects id param in URL
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	taskID := (mux.Vars(r))["id"]
	//taskID is a string, we need to turn it into an id:
	id, _ := primitive.ObjectIDFromHex(taskID)
	// construct the filter for the query:
	filter := bson.M{"_id": id}
	result := collection.FindOne(context.Background(), filter)
	json.NewEncoder(w).Encode(result)
}

// CreateTask : for POST method
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask tasks.Task
	//err := json.NewDecoder(r.Body).Decode(&newTask)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Decoder error:")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("RequestBody: ", reqBody)
	json.Unmarshal(reqBody, &newTask)
	fmt.Println("newTask: ", newTask.Item)

	// insert newTask into database:
	fmt.Println("CreateTask(): ")
	fmt.Println("newTask: ", newTask, "r.Body: ", r.Body, "r.Header", r.Header)
	result, err := collection.InsertOne(context.Background(), newTask)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new record", result.InsertedID)
	fmt.Println("Record: ", *result)

	json.NewEncoder(w).Encode(newTask)
}
