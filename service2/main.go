package main

import (
	"http/httpclient"
	"log"
)

const (
	BASE_URL = "http://localhost:4000/users"
)
	
	// postUser()
	func main() {
		// init client

		client := httpclient.NewHttpClient(BASE_URL)
	
		// proses mengirim request dengan method GET
		resp, err := client.Get("/get")
		if err != nil {
			panic(err)
		}
	
		log.Println(string(resp))
	
		reqData := map[string]interface{}{
			"name": "Reymond - Service 2",
		}
		// proses mengirim request dengan method POST
		resp, err = client.Post("/add", reqData)
		if err != nil {
			panic(err)
		}
	
		log.Println(string(resp))
	}

