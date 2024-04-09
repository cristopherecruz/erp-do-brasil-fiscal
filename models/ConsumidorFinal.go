package models

import "fmt"

type ConsumidorFinal struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	NAO = ConsumidorFinal{"0", "Não"}
	SIM = ConsumidorFinal{"1", "Sim"}
)

// Função para recuperar o ConsumidorFinal com base no código
func RecuperarConsumidorFinalPorCodigo(codigo string) (ConsumidorFinal, error) {
	consumidores := map[string]ConsumidorFinal{
		NAO.Codigo: NAO,
		SIM.Codigo: SIM,
	}

	if consumidor, ok := consumidores[codigo]; ok {
		return consumidor, nil
	}
	return ConsumidorFinal{}, fmt.Errorf("Consumidor final não encontrado para o código %s", codigo)
}
