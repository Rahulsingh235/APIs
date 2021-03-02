package models

import (
	"database/sql"
	"encoding/json"
	"goworkspace/entities"
	"net/http"
)

type ClassModel struct {
	Db *sql.DB
}

func (classModel ClassModel) FindAll() (class []entities.Class, err error) {
	rows, err := classModel.Db.Query("select * from class")
	if err != nil {
		return nil, err
	} else {
		var classes []entities.Class
		for rows.Next() {
			var id int64
			var name string
			var roll int64
			var age int64
			var email_id string
			err2 := rows.Scan(&id, &name, &roll, &age, &email_id)
			if err2 != nil {
				return nil, err2
			} else {
				class := entities.Class{
					Id:    id,
					Name:  name,
					Roll:  roll,
					Age:   age,
					Email: email_id,
				}
				classes = append(classes, class)
			}
		}
		return classes, nil
	}
}

func (classModel ClassModel) Addingstudent(class *entities.Class) (err error) {
	result, err := classModel.Db.Exec("insert into class(id, name, roll, age, email_id) values(?,?,?,?,?)", class.Id, class.Name, class.Roll, class.Age, class.Email)
	if err != nil {
		return err
	} else {
		result.RowsAffected()
		return nil
	}
}
func (classModel ClassModel) Updateparticularemployee(nid int, class *entities.Class) (err error) {
	result, err := classModel.Db.Exec("update class set id = ? where id = ?", class.Id, nid)
	if err != nil {
		return err
	} else {
		result.RowsAffected()
		return nil
	}
}
func (classModel ClassModel) Deletemployee(nid int) error {
	result, err := classModel.Db.Exec("delete from class where id = ?", nid)
	if err != nil {
		return err
	} else {
		result.RowsAffected()
		return nil
	}
}

//func (classModel ClassModel) Up(nid int, class *entities.Class) (err error) {
//	return nil
//}
func (classModel ClassModel) Updateemployee(nid int, class *entities.Class) (err error) {
	result, err := classModel.Db.Exec("update class set id = ?, name = ?, roll = ?, age = ?, email_id = ? where id = ?", class.Id, class.Name, class.Roll, class.Age, class.Email, nid)
	if err != nil {
		return err
	} else {
		result.RowsAffected()
		return nil
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
