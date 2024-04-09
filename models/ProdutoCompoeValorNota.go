package models

import "errors"

type ProdutoCompoeValorNota struct {
	Codigo    string
	Descricao string
}

// Definições possíveis para o campo IndTot
var (
	CompoeValorNota    = ProdutoCompoeValorNota{"1", "Compoe"}
	NaoCompoeValorNota = ProdutoCompoeValorNota{"0", "Não compoe"}
)

// ValorDeCodigo retorna o ProdutoCompoeValorNota correspondente ao código fornecido
func ValorDeCodigoProdutoCompoeValorNota(codigo string) (ProdutoCompoeValorNota, error) {
	switch codigo {
	case CompoeValorNota.Codigo:
		return CompoeValorNota, nil
	case NaoCompoeValorNota.Codigo:
		return NaoCompoeValorNota, nil
	default:
		return ProdutoCompoeValorNota{}, errors.New("Produto Compoe Valor Nota não encontrado!")
	}
}
