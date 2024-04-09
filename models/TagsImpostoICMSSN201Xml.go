package models

type TagsImpostoICMSSN201Xml struct {
	Origem      OrigemNotaFiscal   `xml:"orig"`
	CSOSN       TributacaoICMS     `xml:"CSOSN"`
	ModBCST     ModalidadeBCICMSST `xml:"modBCST"`
	PMVAST      string             `xml:"pMVAST"`
	PRedBCST    string             `xml:"pRedBCST"`
	VBCST       string             `xml:"vBCST"`
	PICMSST     string             `xml:"pICMSST"`
	VICMSST     string             `xml:"vICMSST"`
	PCredSN     string             `xml:"pCredSN"`
	VCredICMSSN string             `xml:"vCredICMSSN"`
}
