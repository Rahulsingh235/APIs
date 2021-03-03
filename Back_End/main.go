package main

import (
	"Back_End/apis/sso_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/apis/employee/findemail", sso_api.FindEmail).Methods("GET")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
