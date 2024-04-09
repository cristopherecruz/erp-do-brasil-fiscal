package models

import "errors"

type NFIndicadorIEDestinatario struct {
	Codigo    string
	Descricao string
}

// Definições possíveis para o indicador de inscrição estadual do destinatário
var (
	ContribuinteICMS                             = NFIndicadorIEDestinatario{"1", "Contribuinte ICMS"}
	ContribuinteIsentoInscricaoContribuintesICMS = NFIndicadorIEDestinatario{"2", "Contribuinte isento inscrição contribuintes ICMS"}
	NaoContribuinte                              = NFIndicadorIEDestinatario{"9", "Não contribuinte"}
)

// ValorDeCodigo retorna o NFIndicadorIEDestinatario correspondente ao código fornecido
func ValorDeCodigo(codigo string) (NFIndicadorIEDestinatario, error) {
	switch codigo {
	case ContribuinteICMS.Codigo:
		return ContribuinteICMS, nil
	case ContribuinteIsentoInscricaoContribuintesICMS.Codigo:
		return ContribuinteIsentoInscricaoContribuintesICMS, nil
	case NaoContribuinte.Codigo:
		return NaoContribuinte, nil
	default:
		return NFIndicadorIEDestinatario{}, errors.New("NF Indicador IE Destinatario não encontrado!")
	}
}
