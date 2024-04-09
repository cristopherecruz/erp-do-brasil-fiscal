package models

type TagsImpostoICMSSN500Xml struct {
	Origem          OrigemNotaFiscal `xml:"orig"`
	CSOSN           TributacaoICMS   `xml:"CSOSN"`
	VBCSTRet        string           `xml:"vBCSTRet,omitempty"`
	PST             string           `xml:"pST,omitempty"`
	VICMSSubstituto string           `xml:"vICMSSubstituto,omitempty"`
	VICMSSTRet      string           `xml:"vICMSSTRet,omitempty"`
}
