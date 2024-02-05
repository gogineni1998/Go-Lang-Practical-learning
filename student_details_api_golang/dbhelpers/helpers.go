package dbhelpers

import (
	"context"
	"fmt"

	"github.com/gogineni1998/student_details_api_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbname = "students"
const collectionname = "marks"

var collection *mongo.Collection
var connectionString = "mongodb+srv://gogineni1998:gf9YL3SJJYcTEgDc@studentcoursesdb.t33vz7u.mongodb.net/?retryWrites=true&w=majority"

func init() {
	clientoptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientoptions)

	if err != nil {
		panic(err)
	}

	collection = client.Database(dbname).Collection(collectionname)
	fmt.Println("collection")
}

func InsertRecord(student *models.Student) *mongo.InsertOneResult {
	response, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		panic(err)
	}

	return response
}

func UpdateRecord(id string, student *models.Student) string {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	updatedResult, err := collection.InsertOne(context.Background(), student)

	if err != nil {
		panic(err)
	}
	return updatedResult.InsertedID.(primitive.ObjectID).String()
}

func DeleteRecord(id string) int64 {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	deleteResult, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		panic(err)
	}
	return deleteResult.DeletedCount
}
