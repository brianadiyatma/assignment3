package model

type StatusReport struct {
	Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	} `json:"status"`
}
