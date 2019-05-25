package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/sayhello", echoPayload)
	log.Println("Go Backend: { HTTPVersion = 1 }; serving on https://localhost:9191/hello/sayhello")
	log.Fatal(http.ListenAndServeTLS(":9191", "../cert/server.crt", "../cert/server.key", nil))
}

func echoPayload(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request connection: %s, path: %s", r.Proto, r.URL.Path[1:])
	defer r.Body.Close()
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Ops! Failed reading body of request. \n %v", err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
