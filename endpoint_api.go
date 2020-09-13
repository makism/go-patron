package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type APIResponse struct {
	Message string  `json:"message"`
}

func PingApi() {
	url := "http://0.0.0.0:5500/v1/api/ping/"
	fmt.Println("Probing", url)

	ctx := context.Background()
	var schemaData = []byte(`{
		"$id": "https://qri.io/schema/",
		"$comment" : "",
		"title": "APIResponse",
		"type": "object",
		"properties": {
			"message": {
				"type": "string"
			}
		},
		"required": ["message"]
	}`)

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	var valid = []byte(`{
	    "message" : "pong"
    }`)

	errs, err := rs.ValidateBytes(ctx, valid)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}

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

	io.Copy(os.Stdout, resp.Body)

	responseMessage := APIResponse{}
	json.NewDecoder(resp.Body).Decode(&responseMessage)
	fmt.Println(responseMessage)
}
