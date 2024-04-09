package models

type TagsCOFINSOutrasOperacoesXml struct {
	SituacaoTributaria   SituacaoTributariaCOFINS `xml:"CST"`
	ValorBaseCalculo     string                   `xml:"vBC"`
	PercentualAliquota   string                   `xml:"pCOFINS"`
	ValorTributo         string                   `xml:"vCOFINS"`
	QuantidadeVendida    string                   `xml:"qBCProd,omitempty"`
	ValorEmReaisAliquota string                   `xml:"vAliqProd,omitempty"`
}
