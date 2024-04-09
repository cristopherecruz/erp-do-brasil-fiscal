package models

import "fmt"

type ModeloNF struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	NFE  = ModeloNF{"55", "NF-e"}
	NFCE = ModeloNF{"65", "NFC-e"}
)

// Função para recuperar o ModeloNF com base no código
func RecuperarModeloNFPorCodigo(codigo string) (ModeloNF, error) {
	modelos := map[string]ModeloNF{
		NFE.Codigo:  NFE,
		NFCE.Codigo: NFCE,
	}

	if m, ok := modelos[codigo]; ok {
		return m, nil
	}
	return ModeloNF{}, fmt.Errorf("Modelo NF não encontrado para o código %s", codigo)
}
