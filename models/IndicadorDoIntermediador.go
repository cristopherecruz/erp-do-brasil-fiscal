package models

import "fmt"

type IndicadorDoIntermediador struct {
	Codigo    string
	Descricao string
}

// Função de construtor para criar uma nova instância de IndicadorDoIntermediador
func NewIndicadorDoIntermediador(codigo, descricao string) IndicadorDoIntermediador {
	return IndicadorDoIntermediador{Codigo: codigo, Descricao: descricao}
}

// Mapa para mapear códigos de indicadores para suas instâncias correspondentes
var indicadorMap = map[string]IndicadorDoIntermediador{
	"0": NewIndicadorDoIntermediador("0", "Operação sem intermediador (em site ou plataforma própria)"),
	"1": NewIndicadorDoIntermediador("1", "Operação em site ou plataforma de terceiros (intermediadores/marketplace)"),
}

// Função para obter uma instância de IndicadorDoIntermediador pelo código
func GetIndicadorDoIntermediadorByCodigo(codigo string) (IndicadorDoIntermediador, error) {
	indicador, ok := indicadorMap[codigo]
	if !ok {
		return IndicadorDoIntermediador{}, fmt.Errorf("IndicadorDoIntermediador com código %s não encontrado", codigo)
	}
	return indicador, nil
}
