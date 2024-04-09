package models

type TagsICMSUFDestinoXml struct {
	ValorBaseCalculoUFDestino    string `xml:"vBCUFDest"`
	ValorBaseCalculoFCPUFDestino string `xml:"vBCFCPUFDest"`
	PercentualFCPUFDestino       string `xml:"pFCPUFDest"`
	PercentualICMSUFDestino      string `xml:"pICMSUFDest"`
	PercentualICMSInterEstadual  string `xml:"pICMSInter"`
	PercentualICMSPartilha       string `xml:"pICMSInterPart"`
	ValorFCPUFDestino            string `xml:"vFCPUFDest"`
	ValorICMSUFDestino           string `xml:"vICMSUFDest"`
	ValorICMSUFRemetente         string `xml:"vICMSUFRemet"`
}
