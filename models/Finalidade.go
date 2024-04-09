package models

import "fmt"

type Finalidade struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	NORMAL               = Finalidade{"1", "Normal"}
	COMPLEMENTAR         = Finalidade{"2", "Complementar"}
	AJUSTE               = Finalidade{"3", "Ajuste"}
	DEVOLUCAO_OU_RETORNO = Finalidade{"4", "Devolução ou retorno"}
)

// Função para recuperar a Finalidade com base no código
func RecuperarFinalidadePorCodigo(codigo string) (Finalidade, error) {
	finalidades := map[string]Finalidade{
		NORMAL.Codigo:               NORMAL,
		COMPLEMENTAR.Codigo:         COMPLEMENTAR,
		AJUSTE.Codigo:               AJUSTE,
		DEVOLUCAO_OU_RETORNO.Codigo: DEVOLUCAO_OU_RETORNO,
	}

	if finalidade, ok := finalidades[codigo]; ok {
		return finalidade, nil
	}
	return Finalidade{}, fmt.Errorf("Finalidade não encontrada para o código %s", codigo)
}
