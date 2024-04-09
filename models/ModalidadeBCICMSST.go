package models

import "errors"

type ModalidadeBCICMSST string

const (
	PrecoTabeladoMaximo ModalidadeBCICMSST = "0"
	ListaNegativa       ModalidadeBCICMSST = "1"
	ListaPositiva       ModalidadeBCICMSST = "2"
	ListaNeutra         ModalidadeBCICMSST = "3"
	MargemValorAgregado ModalidadeBCICMSST = "4"
	Pauta               ModalidadeBCICMSST = "5"
)

var modalidadeMap = map[ModalidadeBCICMSST]string{
	PrecoTabeladoMaximo: "Preço tabelado ou máximo sugerido",
	ListaNegativa:       "Lista Negativa",
	ListaPositiva:       "Lista Positiva",
	ListaNeutra:         "Lista Neutra",
	MargemValorAgregado: "Margem Valor Agregado",
	Pauta:               "Pauta",
}

func (m ModalidadeBCICMSST) Descricao() string {
	return modalidadeMap[m]
}

func ModalidadeBCICMSSTFromCodigo(codigo string) (ModalidadeBCICMSST, error) {
	for modalidade, cod := range modalidadeMap {
		if cod == codigo {
			return modalidade, nil
		}
	}
	return "", errors.New("ModalidadeBCICMSSTNotFound")
}
