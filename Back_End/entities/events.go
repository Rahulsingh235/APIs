package entities

import (
	_ "fmt"
	_ "image"
	_ "image/jpeg"
	"time"
)

type Events struct {
	Event_id      int       `json:"Event_id"`
	Event_name    string    `json:"Event_name"`
	Start_date    time.Time `json:"Start_date"`
	End_date      time.Time `json:"End_date"`
	Image_details string    `json:"Image_details"`
}
