package models

import "fmt"

type TipoNota struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	ENTRADA = TipoNota{"0", "Entrada"}
	SAIDA   = TipoNota{"1", "Saída"}
)

// Função para recuperar o TipoNota com base no código
func RecuperarTipoNotaPorCodigo(codigo string) (TipoNota, error) {
	tipos := map[string]TipoNota{
		ENTRADA.Codigo: ENTRADA,
		SAIDA.Codigo:   SAIDA,
	}

	if t, ok := tipos[codigo]; ok {
		return t, nil
	}
	return TipoNota{}, fmt.Errorf("Tipo Nota não encontrado para o código %s", codigo)
}
