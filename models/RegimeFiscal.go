package models

import (
	"errors"
	"strings"
)

// RegimeFiscal representa os diferentes regimes fiscais
type RegimeFiscal struct {
	Nome      string
	Codigo    string
	Descricao string
}

// Definindo as constantes para os regimes fiscais
var (
	SIMPLES_NACIONAL                 = RegimeFiscal{"Simples Nacional", "1", "Simples nacional"}
	SIMPLES_NACIONAL_EXCESSO_RECEITA = RegimeFiscal{"Simples Nacional excesso de receita", "2", "Simples nacional com excesso de sublimite da receita bruta"}
	NORMAL_FISCAL                    = RegimeFiscal{"Normal", "3", "Regime normal"}
	MEI                              = RegimeFiscal{"MEI", "4", "Micro empreendedor individual"}
	LUCRO_PRESUMIDO                  = RegimeFiscal{"Lucro Presumido", "5", "Lucro Presumido"}
	LUCRO_REAL                       = RegimeFiscal{"Lucro Real", "6", "Lucro real"}
)

// Mapa que mapeia o nome do regime fiscal para sua struct correspondente
var regimeFiscalMap = map[string]RegimeFiscal{
	strings.ToUpper(SIMPLES_NACIONAL.Nome):                 SIMPLES_NACIONAL,
	strings.ToUpper(SIMPLES_NACIONAL_EXCESSO_RECEITA.Nome): SIMPLES_NACIONAL_EXCESSO_RECEITA,
	strings.ToUpper(NORMAL_FISCAL.Nome):                    NORMAL_FISCAL,
	strings.ToUpper(MEI.Nome):                              MEI,
	strings.ToUpper(LUCRO_PRESUMIDO.Nome):                  LUCRO_PRESUMIDO,
	strings.ToUpper(LUCRO_REAL.Nome):                       LUCRO_REAL,
}

// Retorna o RegimeFiscal correspondente ao nome fornecido
func valueOfName(name string) (RegimeFiscal, error) {
	regime, ok := regimeFiscalMap[strings.ToUpper(name)]
	if !ok {
		return RegimeFiscal{}, errors.New("Regime Fiscal não encontrado!")
	}
	return regime, nil
}

// Retorna o RegimeFiscal correspondente ao código fornecido
func valueOfCodigo(codigo string) (RegimeFiscal, error) {
	for _, regime := range regimeFiscalMap {
		if regime.Codigo == codigo {
			return regime, nil
		}
	}
	return RegimeFiscal{}, errors.New("Regime Fiscal não encontrado!")
}
