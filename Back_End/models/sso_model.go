package models

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type EmployeeModel struct {
	Db *sql.DB
}

func (employeeModel EmployeeModel) FindEmail() (email []string, err error) {
	rows, err := employeeModel.Db.Query("select email_id from employee")
	if err != nil {
		return nil, err
	} else {
		var emails []string
		for rows.Next() {
			var email string
			err2 := rows.Scan(&email)
			if err2 != nil {
				return nil, err2
			}
			emails = append(emails, email)
		}

		return emails, nil
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
