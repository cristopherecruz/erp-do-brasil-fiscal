package models

type NFNotaInfoItemImpostoPISAliquota struct {
	SituacaoTributaria SituacaoTributariaPIS `xml:"CST"`
	ValorBaseCalculo   string                `xml:"vBC"`
	PercentualAliquota string                `xml:"pPIS"`
	ValorTributo       string                `xml:"vPIS"`
}
