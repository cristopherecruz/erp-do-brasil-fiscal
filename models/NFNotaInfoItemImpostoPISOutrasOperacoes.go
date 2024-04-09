package models

type NFNotaInfoItemImpostoPISOutrasOperacoes struct {
	SituacaoTributaria   SituacaoTributariaPIS `xml:"CST"`
	ValorBaseCalculo     string                `xml:"vBC"`
	PercentualAliquota   string                `xml:"pPIS"`
	ValorTributo         string                `xml:"vPIS"`
	QuantidadeVendida    string                `xml:"qBCProd,omitempty"`
	ValorEmReaisAliquota string                `xml:"vAliqProd,omitempty"`
}
