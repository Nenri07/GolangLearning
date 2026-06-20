package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nenri07/mongodb/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "ConnectionString"
const dbName = "FirstGoDb"
const colName = "video"

var collection *mongo.Collection

func init() {
	Coptions := options.Client().ApplyURI(connectionString)

	client, error := mongo.Connect(Coptions)
	checkErr(error)
	fmt.Println("Connection built with client")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/// ----- MONGO HELPER METHODS -----

func insertData(video models.Youtube) {
	// Fixed: context.Background()
	Data, err := collection.InsertOne(context.Background(), video)
	checkErr(err)
	fmt.Println("Inserted Data", Data)
}

func updateOne(videoId string) error {
	id, err := bson.ObjectIDFromHex(videoId)
	if err != nil {
		return fmt.Errorf("invalid id %q: %w", videoId, err)
	}

	filter := bson.M{"_id": id}
	update := bson.A{
		bson.M{"$set": bson.M{"isliked": bson.M{"$not": "$isliked"}}},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("matched: %d, modified: %d\n", result.MatchedCount, result.ModifiedCount)
	if result.MatchedCount == 0 {
		fmt.Println("no document matched id:", videoId)
	}
	return nil
}

func deleteOne(videoId string) {
	id, _ := bson.ObjectIDFromHex(videoId)
	filter := bson.M{"_id": id}
	data, err := collection.DeleteOne(context.Background(), filter)
	checkErr(err)

	fmt.Println("This is deleted Data", data)
}

func deleteAll() {
	filter := bson.D{{}}
	data, err := collection.DeleteMany(context.Background(), filter)
	checkErr(err)

	fmt.Println("This is all deleted Data", data)
}

// Fixed return type: Now correctly returns a slice []bson.M
func getAllData() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkErr(err)
	defer cursor.Close(context.Background())

	var movies []bson.M
	// Fixed: cursor.Next() capital letter
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkErr(err)

		movies = append(movies, movie)
	}

	return movies
}

/// ----- CONTROLLER METHODS -----

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alldata := getAllData()
	if alldata != nil {
		json.NewEncoder(w).Encode(alldata)
	} else {
		json.NewEncoder(w).Encode("at this time data is empty")
	}
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie models.Youtube
	// Fixed: json.NewDecoder()
	err := json.NewDecoder(r.Body).Decode(&movie)
	checkErr(err)
	insertData(movie)

	json.NewEncoder(w).Encode(movie)
}

func WatchedSwitch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	params := mux.Vars(r)

	if err := updateOne(params["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOne(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deleteAll()
	// Fixed: Returning a status message instead of missing params reference
	json.NewEncoder(w).Encode("All records successfully deleted.")
}
