package main

import (
	"encoding/json"
	//"github.com/qri-io/jsonschema"
	"log"
	"net/http"
	"time"
	"fmt"
)

type APIResponse struct {
	Message string  `json:"message"`
}

func pingApi() {
	url := "http://0.0.0.0:5500/v1/api/ping/"

	fmt.Println(url)

	client := http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	responseMessage := APIResponse{}
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	fmt.Println(responseMessage)
}
