package models

type TagsImpostoICMS00Xml struct {
	Origem OrigemNotaFiscal `xml:"orig"`
	CST    TributacaoICMS   `xml:"CST"`
	ModBC  ModalidadeBCICMS `xml:"modBC"`
	VBC    string           `xml:"vBC"`
	PICMS  string           `xml:"pICMS"`
	VICMS  string           `xml:"vICMS"`
	PFCP   string           `xml:"pFCP,omitempty"`
	VFCP   string           `xml:"vFCP,omitempty"`
}
