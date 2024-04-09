package models

import "fmt"

type LocalDestinoOperacao struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	OPERACAO_INTERNA       = LocalDestinoOperacao{"1", "Operação interna"}
	OPERACAO_INTERESTADUAL = LocalDestinoOperacao{"2", "Operação interestadual"}
	OPERACAO_COM_EXTERIOR  = LocalDestinoOperacao{"3", "Operação com o exterior"}
)

// Função para recuperar o LocalDestinoOperacao com base no código
func RecuperarLocalDestinoOperacaoPorCodigo(codigo string) (LocalDestinoOperacao, error) {
	operacoes := map[string]LocalDestinoOperacao{
		OPERACAO_INTERNA.Codigo:       OPERACAO_INTERNA,
		OPERACAO_INTERESTADUAL.Codigo: OPERACAO_INTERESTADUAL,
		OPERACAO_COM_EXTERIOR.Codigo:  OPERACAO_COM_EXTERIOR,
	}

	if op, ok := operacoes[codigo]; ok {
		return op, nil
	}
	return LocalDestinoOperacao{}, fmt.Errorf("Local destino operação não encontrado para o código %s", codigo)
}
