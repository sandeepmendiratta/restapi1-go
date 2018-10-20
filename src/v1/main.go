package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func main() {
var router = mux.NewRouter()
router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
router.HandleFunc("/message", handleQryMessage).Methods("GET")
router.HandleFunc("/m/{msg}", handleUrlMessage).Methods("GET")

fmt.Println("Running server!")
log.Fatal(http.ListenAndServe(":3000", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}

//curl http://localhost:3000/message\?msg\=hello
func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
//curl http://localhost:3000/m/helloworld

func handleUrlMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}