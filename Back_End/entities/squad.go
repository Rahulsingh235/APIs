package entities

import _ "fmt"

type Squad struct {
	Squad_names         string  `json:"Squad_names"`
	Overall_scores      float64 `json:"Overall_scores"`
	Present_years_score float64 `json:"Present_years_score"`
	Percentile          float32 `json:"Percentile"`
}
