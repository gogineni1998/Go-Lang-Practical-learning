package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gogineni1998/Go-Lang-Practical-learning/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = "mongodb+srv://gogineni1998:gf9YL3SJJYcTEgDc@studentcoursesdb.t33vz7u.mongodb.net/?retryWrites=true&w=majority"
var dbName = "Netflix"
var collectionName = "watchlist"
var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Connection Established")
}

func insertOneMovie(movie model.Netflix) {
	fmt.Println(movie)
	dbResponse, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		panic(err)
	}
	fmt.Println(dbResponse)
}

func updateOneMovie(movieId string) {
	fmt.Println(movieId, 1)
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		panic(err)
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.M{"watched": true}}}

	dbResponse, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		panic(err)
	}
	fmt.Println(dbResponse.MatchedCount)
}

func deleteOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": id}
	dbResponse, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Println(dbResponse.DeletedCount)
}

func deleteAllMovies() {
	filter := bson.M{}

	dbResponse, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	fmt.Println(dbResponse.DeletedCount)
}

func getAllMovies() []model.Netflix {
	cursor, err := collection.Find(context.Background(), bson.M{})
	var movies []model.Netflix
	if err != nil {
		panic(err)
	}
	cursor.All(context.Background(), &movies)
	return movies
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie *model.Netflix
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &movie)
	if err != nil {
		fmt.Println(err)
	}
	insertOneMovie(*movie)
	w.Write([]byte("Movie add to the watch list"))
}

func MarkMovieWatched(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	w.Write([]byte("Movie Marked as Watched"))
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	w.Write([]byte("movie deleted"))
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	deleteAllMovies()
	w.Write([]byte("watch list is cleared"))
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := getAllMovies()
	data, err := json.Marshal(movies)
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
