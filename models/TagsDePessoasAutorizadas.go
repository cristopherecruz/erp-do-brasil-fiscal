package models

type TagsDePessoasAutorizadas struct {
	CNPJ string `xml:"CNPJ,omitempty"`
	CPF  string `xml:"CPF,omitempty"`
}
