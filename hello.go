package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting the hello-go")

	/**
	* Handle the / route. It'll return hello world as the
	* http body.
	 */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received on / route")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<center><pre>hello world ... !</pre></center>")
	})

	/**
	* Start the http server and listen on the port
	* 8080 for the incoming request to the application.
	 */
	log.Println("Serving content on :8080")
	http.ListenAndServe("0.0.0.0:8080", nil)

}
