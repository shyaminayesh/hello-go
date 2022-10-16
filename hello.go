package main

import (
	"fmt"
	"net/http"
)

func main() {

	/**
	* Handle the / route. It'll return hello world as the
	* http body.
	 */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<center><pre>hello world ... !</pre></center>")
	})

	/**
	* Start the http server and listen on the port
	* 8080 for the incoming request to the application.
	 */
	http.ListenAndServe("localhost:8080", nil)

}
