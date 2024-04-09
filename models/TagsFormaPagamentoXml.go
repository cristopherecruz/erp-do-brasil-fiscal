package models

type TagsFormaPagamentoXml struct {
	DetalhamentoFormasPagamento []TagsDetalharFormaPagamentoXml `xml:"detPag,omitempty"`
	ValorTroco                  string                          `xml:"vTroco,omitempty"`
}
