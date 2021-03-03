package entities

import (
	_ "fmt"
)

type Subposts struct {
	Subposts_id    int    `json:"Subposts_id"`
	Event_id       string `json:"Event_id"`
	Subpost_type   string `json:"Subpost_type"`
	Subpost_detail string `json:"Subpost_detail"`
}
