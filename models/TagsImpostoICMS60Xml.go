package models

type TagsImpostoICMS60Xml struct {
	Origem          OrigemNotaFiscal `xml:"orig"`
	CST             TributacaoICMS   `xml:"CST"`
	VBCSTRet        string           `xml:"vBCSTRet,omitempty"`
	PST             string           `xml:"pST,omitempty"`
	VICMSSubstituto string           `xml:"vICMSSubstituto,omitempty"`
	VICMSSTRet      string           `xml:"vICMSSTRet,omitempty"`
	VBCFCPSTRet     string           `xml:"vBCFCPSTRet,omitempty"`
	PFCPSTRet       string           `xml:"pFCPSTRet,omitempty"`
	VFCPSTRet       string           `xml:"vFCPSTRet,omitempty"`
	PRedBCEfet      string           `xml:"pRedBCEfet,omitempty"`
	VBCEfet         string           `xml:"vBCEfet,omitempty"`
	PICMSEfet       string           `xml:"pICMSEfet,omitempty"`
	VICMSEfet       string           `xml:"vICMSEfet,omitempty"`
}
