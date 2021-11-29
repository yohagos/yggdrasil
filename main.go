package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := NewRouter()

	log.Println("starting server...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(CORS)

	router.HandleFunc("/name", giveName).Methods("POST")

	return router
}

func giveName(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("name")

	if text == "yosef" {
		x, _ := json.Marshal("you did it!!")
		w.Write(x)
	} else {
		x, _ := json.Marshal("you fucked up!!!")
		w.Write(x)
	}
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		//w.Header().Set("Access-Control-Exposed-Headers", "Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)

	})
}
