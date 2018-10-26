package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// To instantiate the gorilla/mux router
	r := mux.NewRouter()

	// On the default page, we will simply serve our static index page
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// To make sure that the API is up and running
	r.Handle("/status", NotImplemented).Methods("GET")

	// To retrieve a list of products that the user can leave feedback on
	r.Handle("/products", NotImplemented).Methods("GET")

	// To capture user feedback on products
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	// To setup the server so one can serve static assets like images, css
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("/static"))))

	// To declare the port and pass in the router
	http.ListenAndServe(":3000", r)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
