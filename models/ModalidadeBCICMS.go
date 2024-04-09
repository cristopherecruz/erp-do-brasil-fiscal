package models

import "errors"

type ModalidadeBCICMS struct {
	Codigo    string
	Descricao string
}

// Definições possíveis para ModalidadeBCICMS
var (
	MVA                   = ModalidadeBCICMS{"0", "Margem de valor agregado"}
	PAUTA                 = ModalidadeBCICMS{"1", "Pauta"}
	PRECO_TABELADO_MAXIMO = ModalidadeBCICMS{"2", "Preço tabelado máximo"}
	VALOR_OPERACAO        = ModalidadeBCICMS{"3", "Valor da operação"}
)

// ValorDeCodigo retorna a ModalidadeBCICMS correspondente ao código fornecido
func ValorDeCodigoModalidade(codigo string) (ModalidadeBCICMS, error) {
	switch codigo {
	case MVA.Codigo:
		return MVA, nil
	case PAUTA.Codigo:
		return PAUTA, nil
	case PRECO_TABELADO_MAXIMO.Codigo:
		return PRECO_TABELADO_MAXIMO, nil
	case VALOR_OPERACAO.Codigo:
		return VALOR_OPERACAO, nil
	default:
		return ModalidadeBCICMS{}, errors.New("Modalidade BCICMS não encontrada")
	}
}
