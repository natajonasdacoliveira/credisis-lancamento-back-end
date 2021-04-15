package main

import (
	"conta-corrente/transferencia/db"
	"conta-corrente/transferencia/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func lancamentoHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {

		debito := jsonMap["debito"].(string)
		var debitoFloat float64
		debitoFloat = 0

		if debito != "" {
			debitoFloat, err = strconv.ParseFloat(debito, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		credito:= jsonMap["credito"].(string)
		var creditoFloat float64
		creditoFloat = 0

		if credito != "" {
			creditoFloat, err = strconv.ParseFloat(credito, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		tipoLancamento:= jsonMap["tipoLancamento"].(string)
		descricao:= jsonMap["descricao"].(string)
		sistema:= jsonMap["sistema"].(string)

		idCooperadoOrigemFK := jsonMap["idCooperadoOrigemFK"].(string)
		var idCooperadoOrigemFKInt uint64
		idCooperadoOrigemFKInt = 0

		if idCooperadoOrigemFK != "" {
			idCooperadoOrigemFKInt, err = strconv.ParseUint(idCooperadoOrigemFK, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		if idCooperadoOrigemFKInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}
		
		idCooperadoDestinoFK := jsonMap["idCooperadoDestinoFK"].(string)
		var idCooperadoDestinoFKInt uint64
		idCooperadoDestinoFKInt = 0

		if idCooperadoDestinoFK != "" {
			idCooperadoDestinoFKInt, err = strconv.ParseUint(idCooperadoDestinoFK, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		if idCooperadoDestinoFKInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		var lancamento model.Lancamento	

		lancamento.Debito = debitoFloat
		lancamento.Credito = creditoFloat
		lancamento.TipoLancamento = tipoLancamento
		lancamento.Descricao = descricao
		lancamento.Sistema = sistema
		lancamento.IDCooperadoOrigem = idCooperadoOrigemFKInt
		lancamento.IDCooperadoDestino = idCooperadoDestinoFKInt

		err :=db.RealizarLancamento(lancamento)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Transferência realizada com sucesso")
	}
}

func cooperadoSaldoHandler(c echo.Context) error {
		jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		id := jsonMap["idCooperado"].(string)
		var idInt uint64
		idInt = 0

		if id != "" {
			idInt, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		if idInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		cooperado, err := db.GetSaldoCooperado(idInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.JSON(http.StatusOK, cooperado)
	}

}

func main() {
	e := echo.New()

	e.POST("/lancamento", lancamentoHandler)
	e.GET("/cooperado/saldo", cooperadoSaldoHandler)

	e.Logger.Fatal(e.Start(":1323"))
}