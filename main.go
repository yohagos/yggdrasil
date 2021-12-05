package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Name struct {
	key   string
	value string
}

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

	router.HandleFunc("/name", giveName)

	return router
}

func giveName(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	//log.Println(body)
	var req interface{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
	}
	log.Println(req)
	test := req.(map[string]interface{})
	fmt.Sprintln("\n%T", test["name"])
	if test["name"] == "yosef" {
		log.Println("did it")
		x, _ := json.Marshal("you did it!!")
		w.Write(x)
	} else {
		log.Println("fucked up")
		x, _ := json.Marshal("you fucked up!!!")
		w.Write(x)
	}
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		//allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
		//methods := "GET, POST, DELETE"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		//w.Header().Set("Access-Control-Exposed-Headers", "Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
