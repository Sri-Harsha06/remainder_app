package controller

import(
	"remainder_app/services"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"remainder_app/model"
	"remainder_app/functions"
	"context"
)

var client *mongo.Client
var event model.Event

func AddEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&event)
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	result, _ := functions.InsertData(collection, event)
	json.NewEncoder(response).Encode(result)
}

func ReadEventById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	err := functions.FindDataById(collection, model.Event{Id: id}).Decode(&event)
	// err := collection.FindOne(context.Background(), Event{Id: id}).Decode(&event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(event)
}

func GetEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	// cursor, err := collection.Find(ctx, bson.M{})
	cursor, err := functions.GetData(collection, model.Event{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func ReadEventByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	params := mux.Vars(request)
	name := params["name"]
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	cursor, err := functions.GetData(collection, model.Event{Name: name})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func ReadEventByEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	params := mux.Vars(request)
	event := params["event"]
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	cursor, err := functions.GetData(collection, model.Event{Event: event})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var event model.Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func ReadEventByDate(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	params := mux.Vars(request)
	date := params["date"]
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	cursor, err := functions.GetData(collection, model.Event{Date: date})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func UpdateEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	_ = json.NewDecoder(request.Body).Decode(&event)
	result, err := functions.UpdateData(collection, event)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)
}

func DeleteEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	client=services.GetInstance()
	collection := client.Database("events").Collection("event")
	result, err := functions.DeleteData(collection, model.Event{Id: id})
	if err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(result)
}