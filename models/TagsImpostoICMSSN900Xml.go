package models

type TagsImpostoICMSSN900Xml struct {
	Origem      OrigemNotaFiscal   `xml:"orig"`
	CSOSN       TributacaoICMS     `xml:"CSOSN"`
	ModBC       ModalidadeBCICMS   `xml:"modBC"`
	VBC         string             `xml:"vBC"`
	PRedBC      string             `xml:"pRedBC,omitempty"`
	PICMS       string             `xml:"pICMS,omitempty"`
	VICMS       string             `xml:"vICMS,omitempty"`
	ModBCST     ModalidadeBCICMSST `xml:"modBCST,omitempty"`
	PMVAST      string             `xml:"pMVAST,omitempty"`
	PRedBCST    string             `xml:"pRedBCST,omitempty"`
	VBCST       string             `xml:"vBCST,omitempty"`
	PICMSST     string             `xml:"pICMSST,omitempty"`
	VICMSST     string             `xml:"vICMSST,omitempty"`
	PCredSN     string             `xml:"pCredSN,omitempty"`
	VCredICMSSN string             `xml:"vCredICMSSN,omitempty"`
}
