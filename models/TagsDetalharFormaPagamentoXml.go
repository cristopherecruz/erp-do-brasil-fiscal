package models

type TagsDetalharFormaPagamentoXml struct {
	IndicadorFormaPagamento IndicadorFormaPagamento `xml:"indPag,omitempty"`
	MeioPagamento           FormaDePagamento        `xml:"tPag"`
	DescricaoDoPagamento    string                  `xml:"xPag,omitempty"`
	ValorPagamento          string                  `xml:"vPag"`
	Cartao                  TagsNotaInfoCartaoXml   `xml:"card,omitempty"`
}
