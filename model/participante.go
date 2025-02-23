package model

type Participante struct {
	ID         int    `json:"id" gorm:"AUTO_INCREMENT:primaryKey"`
	Nome       string `json:"nome"`
	Residencia string `json:"residencia"`
	Ocupacao   string `json:"ocupacao"`
	Status     bool   `json:"status"`
}
