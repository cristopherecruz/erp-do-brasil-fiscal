package models

import "fmt"

type Ambiente struct {
	Codigo    string
	Descricao string
}

// Valores pré-definidos
var (
	PRODUCAO    = Ambiente{"1", "Produção"}
	HOMOLOGACAO = Ambiente{"2", "Homologação"}
)

// Função para recuperar o Ambiente com base na descrição
func RecuperarAmbientePorDescricao(descricao string) (Ambiente, error) {
	ambientes := map[string]Ambiente{
		PRODUCAO.Descricao:    PRODUCAO,
		HOMOLOGACAO.Descricao: HOMOLOGACAO,
	}

	if ambiente, ok := ambientes[descricao]; ok {
		return ambiente, nil
	}
	return Ambiente{}, fmt.Errorf("Ambiente não encontrado para a descrição %s", descricao)
}

// Função para recuperar o Ambiente com base no código
func RecuperarAmbientePorCodigo(codigo string) (Ambiente, error) {
	ambientes := map[string]Ambiente{
		PRODUCAO.Codigo:    PRODUCAO,
		HOMOLOGACAO.Codigo: HOMOLOGACAO,
	}

	if ambiente, ok := ambientes[codigo]; ok {
		return ambiente, nil
	}
	return Ambiente{}, fmt.Errorf("Ambiente não encontrado para o código %s", codigo)
}
