package model

type Cooperado struct {
	IDCooperado uint64  `db:"idCooperado" json:"idCooperado"`
	Conta       string  `db:"conta" json:"conta"`
	Saldo       float64 `db:"saldo" json:"saldo"`
	Nome        string  `db:"nome" json:"nome"`
}