package db

import (
	"conta-corrente/transferencia/model"
	"errors"
	"time"
)

func RealizarLancamento(lancamento model.Lancamento) (err error) {
	db := InitDb()
	
	oldCooperadoOrigem := model.Cooperado{}
	oldCooperadoDestino := model.Cooperado{}
	
	if(lancamento.Debito == 0 && lancamento.Credito == 0) {
		return errors.New("valor inválido")
	}

	err = db.Get(&oldCooperadoOrigem, "SELECT * FROM cooperado WHERE idCooperado= ? LIMIT 1", lancamento.IDCooperadoOrigem)
	if oldCooperadoOrigem.IDCooperado < 1 || err != nil {
		return errors.New("cooperado origem não encontrado")
	}

	if oldCooperadoOrigem.Saldo < lancamento.Debito {
		return errors.New("saldo insuficiente")
	}

	err = db.Get(&oldCooperadoDestino, "SELECT * FROM cooperado WHERE idCooperado= ? LIMIT 1", lancamento.IDCooperadoDestino)
	if oldCooperadoDestino.IDCooperado < 1 || err != nil {
		return errors.New("cooperado destino não encontrado")
	}

	var (
		novoSaldoCooperadoOrigem float64 
		novoSaldoCooperadoDestino float64
	)

	novoSaldoCooperadoOrigem = oldCooperadoOrigem.Saldo - lancamento.Debito
	novoSaldoCooperadoDestino = oldCooperadoDestino.Saldo + lancamento.Debito

	var date = time.Now()
	dateTime := date.Format(time.RFC3339)

	tx := db.MustBegin()

	tx.MustExec("INSERT INTO lancamento(debito, credito, tipoLancamento, descricao, createdAt, sistema, idCooperadoOrigemFK, idCooperadoDestinoFK) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", lancamento.Debito, lancamento.Credito, lancamento.TipoLancamento, lancamento.Descricao, dateTime, lancamento.Sistema, lancamento.IDCooperadoOrigem, lancamento.IDCooperadoDestino)

	tx.MustExec("UPDATE cooperado SET saldo = ? WHERE idCooperado = ?", novoSaldoCooperadoOrigem, oldCooperadoOrigem.IDCooperado)
	tx.MustExec("UPDATE cooperado SET saldo = ? WHERE idCooperado = ?", novoSaldoCooperadoDestino, oldCooperadoDestino.IDCooperado)

	tx.Commit()

	return err
}