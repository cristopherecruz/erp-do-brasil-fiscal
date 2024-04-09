package models

type TagsDetalharProdutoServicoXml struct {
	NItem     int                     `xml:"nItem,attr,omitempty"`
	InfAdProd string                  `xml:"infAdProd,omitempty"`
	Produto   TagsProdutoServicoXml   `xml:"prod"`
	Imposto   TagsIncidenteImpostoXml `xml:"imposto"`
}
