package main

import (
	"fmt"
	"goworkspace/apis/class_api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/apis/class/findall", class_api.FindAll).Methods("GET")
	router.HandleFunc("/apis/class/addemp", class_api.Addingstudent).Methods("POST")
	router.HandleFunc("/apis/class/{id}", class_api.Updateemployee).Methods("PUT")
	router.HandleFunc("/apis/class/{id}", class_api.Updateparticularemployee).Methods("PATCH")
	router.HandleFunc("/apis/class/{id}", class_api.Deletemployee).Methods("DELETE")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

/*[
    {
        "Id": 1,
        "Name": "rahul",
        "Roll": 1,
        "Age": 23,
        "Email": "rahul.singh@nineleaps.com"
    },
    {
        "Id": 2,
        "Name": "apurv",
        "Roll": 2,
        "Age": 24,
        "Email": "apurv.prakash@nineleaps.com"
    },
    {
        "Id": 3,
        "Name": "rajat",
        "Roll": 3,
        "Age": 42,
        "Email": "rajat.singh@nineleaps.com"
    },
    {
        "Id": 4,
        "Name": "srimraj",
        "Roll": 4,
        "Age": 23,
        "Email": "srimraj.biswal@nineleaps.com"
    }
]*/
