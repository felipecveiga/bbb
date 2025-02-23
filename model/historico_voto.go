package model

import (
	"time"
)

type HistoricoVoto struct {
	ID             int       `json:"id" gorm:"AUTO_INCREMENT:primaryKey"`
	IdParticipante int       `json:"id_participante"`
	Ip             string    `json:"ip" gorm:"type:varchar(39);"`
	Created_at     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
