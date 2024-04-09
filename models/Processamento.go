package models

import "fmt"

type Processamento struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	ASSINCRONO = Processamento{"0", "Processamento Assíncrono"}
	SINCRONO   = Processamento{"1", "Processamento Síncrono"}
)

// Função para recuperar o Processamento com base no código
func RecuperarProcessamentoPorCodigo(codigo string) (Processamento, error) {
	processamentos := map[string]Processamento{
		ASSINCRONO.Codigo: ASSINCRONO,
		SINCRONO.Codigo:   SINCRONO,
	}

	if p, ok := processamentos[codigo]; ok {
		return p, nil
	}

	return Processamento{}, fmt.Errorf("Processamento não encontrado para o código %s", codigo)
}
