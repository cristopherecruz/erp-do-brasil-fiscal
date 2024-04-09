package models

type TagsImpostoICMS40Xml struct {
	Origem OrigemNotaFiscal `xml:"orig"`
	CST    TributacaoICMS   `xml:"CST"`
}
