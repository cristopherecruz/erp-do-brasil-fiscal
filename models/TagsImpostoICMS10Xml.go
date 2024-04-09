package models

type TagsImpostoICMS10Xml struct {
	Origem   OrigemNotaFiscal   `xml:"orig"`
	CST      TributacaoICMS     `xml:"CST"`
	ModBC    ModalidadeBCICMS   `xml:"modBC"`
	VBC      string             `xml:"vBC"`
	PICMS    string             `xml:"pICMS"`
	VICMS    string             `xml:"vICMS"`
	VBCFCP   string             `xml:"vBCFCP,omitempty"`
	PFCP     string             `xml:"pFCP,omitempty"`
	VFCP     string             `xml:"vFCP,omitempty"`
	ModBCST  ModalidadeBCICMSST `xml:"modBCST"`
	PMVAST   string             `xml:"pMVAST,omitempty"`
	PRedBCST string             `xml:"pRedBCST,omitempty"`
	VBCST    string             `xml:"vBCST"`
	PICMSST  string             `xml:"pICMSST"`
	VICMSST  string             `xml:"vICMSST"`
	VBCFCPST string             `xml:"vBCFCPST,omitempty"`
	PFCPST   string             `xml:"pFCPST,omitempty"`
	VFCPST   string             `xml:"vFCPST,omitempty"`
}
