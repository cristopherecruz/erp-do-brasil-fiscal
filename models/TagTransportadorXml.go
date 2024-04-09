package models

type TagTransportadorXml struct {
	CNPJ               string `xml:"CNPJ,omitempty"`
	RazaoSocial        string `xml:"xNome,omitempty"`
	InscricaoEstadual  string `xml:"IE,omitempty"`
	Endereco           string `xml:"xEnder,omitempty"`
	DescricaoMunicipio string `xml:"xMun,omitempty"`
	UF                 string `xml:"UF,omitempty"`
}
