package controller

import (
	"context"
	"encoding/json"
	"net/http"
	cd "remainder_App/client_discovery"
	"remainder_app/customlogger"
	"remainder_app/functions"
	"remainder_app/model"
	"remainder_app/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var event model.Event

func AddEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&event)
	client = services.GetInstance()
	collection := client.Database("events").Collection("event")
	result, err := functions.InsertData(collection, event)
	logger := customlogger.GetInstance()
	if err != nil {
		logger.ErrorLogger.Println(err)
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	} else {
		logger.InfoLogger.Println("Inserted object with Id:" + event.Id)
		cd.Client_discovery()
		json.NewEncoder(response).Encode(result)
	}
}

func ReadEventById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	client = services.GetInstance()
	collection := client.Database("events").Collection("event")
	logger := customlogger.GetInstance()
	err := functions.FindDataById(collection, model.Event{Id: id}).Decode(&event)
	// err := collection.FindOne(context.Background(), Event{Id: id}).Decode(&event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.ErrorLogger.Println(err.Error())
		return
	}
	logger.InfoLogger.Println("Fetched event by ID:" + id)
	json.NewEncoder(response).Encode(event)
}

func GetEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	client = services.GetInstance()
	collection := client.Database("events").Collection("event")
	logger := customlogger.GetInstance()
	// cursor, err := collection.Find(ctx, bson.M{})
	cursor, err := functions.GetData(collection, model.Event{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.ErrorLogger.Println(err.Error())
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
		logger.ErrorLogger.Println(err.Error())
		return
	}
	logger.InfoLogger.Println("Fetched all events")
	json.NewEncoder(response).Encode(events)
}

func ReadEventByName(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	params := mux.Vars(request)
	name := params["name"]
	client = services.GetInstance()
	logger := customlogger.GetInstance()
	collection := client.Database("events").Collection("event")
	cursor, err := functions.GetData(collection, model.Event{Name: name})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.ErrorLogger.Println(err.Error())
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
		logger.ErrorLogger.Println(err.Error())
		return
	}
	logger.InfoLogger.Println("Fetched event by name:" + name)
	json.NewEncoder(response).Encode(events)
}

func ReadEventByEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	params := mux.Vars(request)
	event := params["event"]
	client = services.GetInstance()
	logger := customlogger.GetInstance()
	collection := client.Database("events").Collection("event")
	cursor, err := functions.GetData(collection, model.Event{Event: event})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.ErrorLogger.Println(err.Error())
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
		logger.ErrorLogger.Println(err.Error())
		return
	}
	logger.InfoLogger.Println("Fetched event by Event:" + event)
	json.NewEncoder(response).Encode(events)
}

func ReadEventByDate(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []model.Event
	params := mux.Vars(request)
	date := params["date"]
	client = services.GetInstance()
	logger := customlogger.GetInstance()
	collection := client.Database("events").Collection("event")
	cursor, err := functions.GetData(collection, model.Event{Date: date})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.ErrorLogger.Println(err.Error())
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
		logger.ErrorLogger.Println(err.Error())
		return
	}
	logger.InfoLogger.Println("Fetched event by Date:" + date)
	json.NewEncoder(response).Encode(events)
}

func UpdateEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	client = services.GetInstance()
	logger := customlogger.GetInstance()
	collection := client.Database("events").Collection("event")
	_ = json.NewDecoder(request.Body).Decode(&event)
	result, err := functions.UpdateData(collection, event)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		panic(err)
	}
	logger.InfoLogger.Println("Updated Event With Id:" + event.Id)
	json.NewEncoder(response).Encode(result)
}

func DeleteEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	client = services.GetInstance()
	logger := customlogger.GetInstance()
	collection := client.Database("events").Collection("event")
	result, err := functions.DeleteData(collection, model.Event{Id: id})
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		panic(err)
	}
	logger.InfoLogger.Println("Deleted Event With Id:" + id)
	json.NewEncoder(response).Encode(result)
}
