package models

type TagsNotaInfoCartaoXml struct {
	TipoIntegracao                  TipoIntegracaoPagamento `xml:"tpIntegra"`
	Cnpj                            string                  `xml:"CNPJ,omitempty"`
	OperadoraCartao                 OperadoraCartao         `xml:"tBand,omitempty"`
	NumeroAutorizacaoOperacaoCartao string                  `xml:"cAut,omitempty"`
}
