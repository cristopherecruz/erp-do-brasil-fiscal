package models

type TagsCOFINSAliquotaXml struct {
	SituacaoTributaria SituacaoTributariaCOFINS `xml:"CST"`
	ValorBaseCalculo   string                   `xml:"vBC"`
	PercentualAliquota string                   `xml:"pCOFINS"`
	Valor              string                   `xml:"vCOFINS"`
}
