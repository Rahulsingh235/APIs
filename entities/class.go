package entities

import _ "fmt"

type Class struct {
	Id    int64  `json:"Id"`
	Name  string `json:"Name"`
	Roll  int64  `json:"Roll"`
	Age   int64  `json:"Age"`
	Email string `json:"Email"`
}

//func (class Class) ToString() string {
//	return fmt.Sprintf("id: %d\nname: %s\nroll: %d\nage: %d\nemail_id: %s", class.Id, class.Name, class.Roll, class.Age, class.Email)
//}
