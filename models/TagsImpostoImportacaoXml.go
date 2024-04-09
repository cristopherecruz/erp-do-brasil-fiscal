package models

type TagsImpostoImportacaoXml struct {
	ValorBaseCalculo       string `xml:"vBC"`
	ValorDespesaAduaneira  string `xml:"vDespAdu"`
	ValorImpostoImportacao string `xml:"vII"`
	ValorIOF               string `xml:"vIOF"`
}
