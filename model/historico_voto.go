package model

import "time"

type HistoricoVoto struct {
	ID             string    `json:"id"`
	IdParticipante string    `json:"id_participante"`
	Ip             string    `json:"ip"`
	Created_at     time.Time `json:"created_at"`
}
