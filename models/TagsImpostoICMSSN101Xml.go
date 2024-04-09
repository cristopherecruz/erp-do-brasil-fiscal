package models

type TagsImpostoICMSSN101Xml struct {
	Origem      OrigemNotaFiscal `xml:"orig"`
	CSOSN       TributacaoICMS   `xml:"CSOSN"`
	PCredSN     string           `xml:"pCredSN"`
	VCredICMSSN string           `xml:"vCredICMSSN"`
}
