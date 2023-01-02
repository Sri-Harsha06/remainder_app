package main

import (
	// "fmt"
	"net/http"
	"remainder_app/controller"
	"github.com/gorilla/mux"
	"remainder_app/customlogger"
)


func init(){
	
}

func main() {
	// fmt.Println(time.Now().AddDate(0, 0, 1).Format("02-01-2006"))
	// fmt.Println(time.Now().Format("15:04"))
	// fmt.Println("Starting the application...")
	logger := customlogger.GetInstance()
	logger.InfoLogger.Println("Starting")
	router := mux.NewRouter()
	router.HandleFunc("/addevent", controller.AddEvent).Methods("POST")
	router.HandleFunc("/allevents", controller.GetEvents).Methods("GET")
	router.HandleFunc("/event/{id}", controller.ReadEventById).Methods("GET")
	router.HandleFunc("/event2/{name}", controller.ReadEventByName).Methods("GET")
	router.HandleFunc("/event3/{event}", controller.ReadEventByEvent).Methods("GET")
	router.HandleFunc("/event4/{date}", controller.ReadEventByDate).Methods("GET")
	router.HandleFunc("/update", controller.UpdateEvent).Methods("POST")
	router.HandleFunc("/delete/{id}", controller.DeleteEvent).Methods("GET")
	// log.SetOutput(os.Stdout)
	// log.Println("Starting the application...")
	http.ListenAndServe(":12345", router)
}
