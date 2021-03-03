package entities

import _ "fmt"

type Employee struct {
	E_id       string `json:"E_id"`
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
	Email_id   string `json:"Email_id"`
	Gender     string `json:"Gender"`
	College    string `json:"College"`
}
