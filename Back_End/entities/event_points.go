package entities

import _ "fmt"

type Event_points struct {
	Event_id   string  `json:"E_id"`
	Squad_name string  `json:"Squad_name"`
	Points     float64 `json:"Points"`
}
