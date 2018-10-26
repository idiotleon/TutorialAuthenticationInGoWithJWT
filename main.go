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

	// To setup the server so one can serve static assets like images, css
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("/static"))))

	// To declare the port and pass in the router
	http.ListenAndServe(":3000", r)
}
