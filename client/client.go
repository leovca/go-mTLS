package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s", body)
}
