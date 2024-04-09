package models

type TagsImpostoICMSSN102Xml struct {
	Origem OrigemNotaFiscal `xml:"orig"`
	CSOSN  TributacaoICMS   `xml:"CSOSN"`
}
