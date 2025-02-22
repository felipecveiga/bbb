package model

type Participante struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Residencia string `json:"residencia"`
	Ocupacao   string `json:"ocupacao"`
	Status     bool   `json:"status"`
}
