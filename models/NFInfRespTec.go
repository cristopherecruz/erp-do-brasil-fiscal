package models

import "encoding/xml"

type NFInfRespTec struct {
	XMLName  xml.Name `xml:"infRespTec" namespace:"http://www.portalfiscal.inf.br/nfe"`
	CNPJ     string   `xml:"CNPJ"`
	Contato  string   `xml:"xContato"`
	Email    string   `xml:"email"`
	Telefone string   `xml:"fone"`
}
