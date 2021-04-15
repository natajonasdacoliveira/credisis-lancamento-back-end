package model

import "time"

type Lancamento struct {
	IDLancamento       uint64    `db:"idLancamento" json:"idLancamento"`
	Debito             float64   `db:"debito" json:"debito"`
	Credito            float64   `db:"credito" json:"credito"`
	TipoLancamento     string    `db:"tipoLancamento" json:"tipoLancamento"`
	Descricao          string    `db:"descricao" json:"descricao"`
	CreatedAt          time.Time `db:"createdAt" json:"createdAt"`
	Sistema            string    `db:"sistema" json:"sistema"`
	IDCooperadoOrigem  uint64    `db:"idCooperadoOrigemFK" json:"idCooperadoOrigemFK"`
	IDCooperadoDestino uint64    `db:"idCooperadoDestinoFK" json:"idCooperadoDestinoFK"`
}
