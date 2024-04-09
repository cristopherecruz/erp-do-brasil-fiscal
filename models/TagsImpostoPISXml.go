package models

type TagsImpostoPISXml struct {
	Aliquota        NFNotaInfoItemImpostoPISAliquota        `xml:"PISAliq,omitempty"`
	OutrasOperacoes NFNotaInfoItemImpostoPISOutrasOperacoes `xml:"PISOutr,omitempty"`
	PISNT           []TagsImpostoPISNTXml                   `xml:"PISNT,omitempty"`
}
