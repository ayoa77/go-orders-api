package main

import(
	"log"
	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
	"../pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", handlers.GetAllOrders).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode("Hello Orders")
	})

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}