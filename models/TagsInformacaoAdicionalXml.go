package models

type TagsInformacaoAdicionalXml struct {
	InformacoesAdicionaisInteresseFisco            string `xml:"infAdFisco,omitempty"`
	InformacoesComplementaresInteresseContribuinte string `xml:"infCpl,omitempty"`
}
