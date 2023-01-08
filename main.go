package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"remainder_app/controller"
	"remainder_app/customlogger"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"

	// "github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/gorilla/mux"
)

// var configurations RegistrationVariables

// func init() {
// 	registryType := "NETFLIX_EUREKA"
// 	serviceRegistryURL := "http://localhost:8761/eureka/apps/"
// 	configurations = RegistrationVariables{registryType, serviceRegistryURL}
// }

func main() {
	serviceRegistryWithConsul()
	logger := customlogger.GetInstance()
	logger.InfoLogger.Println("Starting")
	// client := eureka.NewClient([]string{
	// 	"http://localhost:8761/eureka/",
	// })
	// fmt.Print(client)
	// applications, _ := client.GetApplications()
	fmt.Println("hii")
	router := mux.NewRouter()
	router.HandleFunc("/addevent", controller.AddEvent).Methods("POST")
	router.HandleFunc("/allevents", controller.GetEvents).Methods("GET")
	router.HandleFunc("/event/{id}", controller.ReadEventById).Methods("GET")
	router.HandleFunc("/event2/{name}", controller.ReadEventByName).Methods("GET")
	router.HandleFunc("/event3/{event}", controller.ReadEventByEvent).Methods("GET")
	router.HandleFunc("/event4/{date}", controller.ReadEventByDate).Methods("GET")
	router.HandleFunc("/update", controller.UpdateEvent).Methods("POST")
	router.HandleFunc("/check", check)
	router.HandleFunc("/helloworld", helloworld)
	router.HandleFunc("/delete/{id}", controller.DeleteEvent).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func serviceRegistryWithConsul() {
	config := consulapi.DefaultConfig()
	fmt.Print(config)
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}
	serviceID := "go_micro_1"
	port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	address := getHostname()

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "go_micro_1",
		Port:    port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	return
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}

func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Consul check")
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	log.Println("helloworld service is called.")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world.")
}
