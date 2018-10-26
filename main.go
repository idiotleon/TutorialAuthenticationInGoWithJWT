package main

import (
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	// To instantiate the gorilla/mux router
	r := mux.NewRouter()

	// On the default page, we will simply serve our static index page
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// To make sure that the API is up and running
	r.Handle("/status", StatusHandler).Methods("GET")

	// To retrieve a list of products that the user can leave feedback on
	r.Handle("/products", jwtMiddleware.Handler(ProductsHandler)).Methods("GET")

	// To capture user feedback on products
	r.Handle("/products/{slug}/feedback", jwtMiddleware.Handler(AddFeedbackHandler)).Methods("POST")

	// To setup the server so one can serve static assets like images, css, js
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	r.Handle("/get-token", GetTokenHandler).Methods("GET")

	// To declare the port and pass in the router
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// To convert the slice of products to JSON
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

// To set a global string for our secret
var mySigningKey = []byte("secret")

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// To create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// To create a map to store claims
	claims := token.Claims.(jwt.MapClaims)

	// To set token claims
	// Hard-coded claims
	claims["admin"] = true
	claims["name"] = "Leon"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// To sign the token with our secret
	tokenString, _ := token.SignedString(mySigningKey)

	// To write the token to the browser window
	w.Write([]byte(tokenString))
})

// To verify the token
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
