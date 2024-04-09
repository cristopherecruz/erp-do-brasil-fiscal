package models

import "encoding/xml"

type TagsEmitenteXml struct {
	XMLName           xml.Name                `xml:"TagsEmitenteXml"`
	CNPJ              string                  `xml:"CNPJ,omitempty"`
	CPF               string                  `xml:"CPF,omitempty"`
	RazaoSocial       string                  `xml:"xNome"`
	NomeFantasia      string                  `xml:"xFant,omitempty"`
	Endereco          TagsEnderecoEmitenteXml `xml:"enderEmit"`
	InscricaoEstadual string                  `xml:"IE"`
	RegimeTributario  RegimeFiscal            `xml:"CRT"`
}
