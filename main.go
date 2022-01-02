package main

import(
	"log"
	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("Hello Orders")
	})

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}