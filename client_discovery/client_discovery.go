package client_discovery

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var url string

func Client_discovery() {
	serviceDiscoveryWithConsul()
	fmt.Println("Starting Client.")
	var client = &http.Client{
		Timeout: time.Second * 30,
	}
	hello(client)
}

func serviceDiscoveryWithConsul() {
	config := consulapi.DefaultConfig()
	consul, error := consulapi.NewClient(config)
	if error != nil {
		fmt.Println(error)
	}
	services, error := consul.Agent().Services()
	if error != nil {
		fmt.Println(error)
	}
	// fmt.Print(services["go_micro_1"])
	service := services["go_micro_2"]
	fmt.Print(service)
	address := service.Address
	port := service.Port
	url = fmt.Sprintf("http://%s:%v/tmrevent", address, port)
	fmt.Print(url)
}

func hello(client *http.Client) {
	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s\n", body)
}
