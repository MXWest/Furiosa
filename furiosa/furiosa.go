package main

import (
	"net/http"
	"html"
	"fmt"
)


func main() {
	fmt.Println("Starting Up.")
	http.HandleFunc("/health", health)
	http.ListenAndServe("localhost:50005", nil)
}


func health(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "received %q", html.EscapeString(r.URL.Path))
}
