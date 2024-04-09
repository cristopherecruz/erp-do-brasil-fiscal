package models

import "encoding/xml"

type TagsEnderecoEmitenteXml struct {
	XMLName            xml.Name `xml:"TagsEnderecoEmitenteXml"`
	Logradouro         string   `xml:"xLgr"`
	Numero             string   `xml:"nro"`
	Complemento        string   `xml:"xCpl,omitempty"`
	Bairro             string   `xml:"xBairro"`
	CodigoMunicipio    string   `xml:"cMun"`
	DescricaoMunicipio string   `xml:"xMun"`
	UF                 string   `xml:"UF"`
	CEP                string   `xml:"CEP,omitempty"`
	CodigoPais         Pais     `xml:"cPais,omitempty"`
	DescricaoPais      string   `xml:"xPais,omitempty"`
	Telefone           string   `xml:"fone,omitempty"`
}
