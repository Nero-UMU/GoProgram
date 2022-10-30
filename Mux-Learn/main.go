package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func msgHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", msgHandle)

	http.ListenAndServe(":8080", r)
}
