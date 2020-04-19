package main

import (
	"net/http"

	lib "./lib"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/deals", lib.GetDeals).Methods("GET")
	router.HandleFunc("/merchants", lib.GetMerchants).Methods("GET")
	router.HandleFunc("/merchant", lib.GetMerchant).Methods("GET")
	http.ListenAndServe(":3000", router)
}
