package main

import (
	"fmt"
	"net/http"
	"remainder_app/controller"
	"remainder_app/customlogger"

	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/gorilla/mux"
)

var configurations RegistrationVariables

func init() {
	registryType := "NETFLIX_EUREKA"
	serviceRegistryURL := "http://localhost:8761/eureka/apps/"
	configurations = RegistrationVariables{registryType, serviceRegistryURL}
}

func main() {

	logger := customlogger.GetInstance()
	logger.InfoLogger.Println("Starting")
	client := eureka.NewClient([]string{
		"http://localhost:8761/eureka/",
	})
	fmt.Print(client)
	applications, _ := client.GetApplications()
	fmt.Println(applications)
	router := mux.NewRouter()
	router.HandleFunc("/addevent", controller.AddEvent).Methods("POST")
	router.HandleFunc("/allevents", controller.GetEvents).Methods("GET")
	router.HandleFunc("/event/{id}", controller.ReadEventById).Methods("GET")
	router.HandleFunc("/event2/{name}", controller.ReadEventByName).Methods("GET")
	router.HandleFunc("/event3/{event}", controller.ReadEventByEvent).Methods("GET")
	router.HandleFunc("/event4/{date}", controller.ReadEventByDate).Methods("GET")
	router.HandleFunc("/update", controller.UpdateEvent).Methods("POST")
	router.HandleFunc("/delete/{id}", controller.DeleteEvent).Methods("GET")
	http.ListenAndServe(":8080", router)
}
