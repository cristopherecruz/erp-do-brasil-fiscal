package models

type TagsImpostoCOFINSXml struct {
	Aliquota           TagsCOFINSAliquotaXml        `xml:"COFINSAliq,omitempty"`
	OutrasOperacoesXml TagsCOFINSOutrasOperacoesXml `xml:"COFINSOutr,omitempty"`
	COFINSNT           []CofinsNaoTributavelXml     `xml:"COFINSNT,omitempty"`
}
