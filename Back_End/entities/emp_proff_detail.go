package entities

import _ "fmt"

/*type Interest struct {
	details map[string]string
}*/

type Emp_proff_detail struct {
	E_id         string            `json:"E_id"`
	Squad_names  string            `json:"Squad_names"`
	Interests    map[string]string `json:"Interests"`
	Years_of_exp float64           `json:"Years_of_exp"`
	Designation  string            `json:"Designation"`
	Project      string            `json:"Project"`
	Role         string            `json:"Role"`
}
