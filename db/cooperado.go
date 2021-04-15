package db

import (
	"conta-corrente/transferencia/model"
	"log"
)


func GetSaldoCooperado(idCooperado uint64) (model.Cooperado, error) {
	db := InitDb()

	cooperado := model.Cooperado{}

	err := db.Get(&cooperado, "SELECT * FROM cooperado WHERE idCooperado = ?", idCooperado)
	if err != nil {
		log.Println(err)
		return model.Cooperado{}, err
	}

	return cooperado, err
}