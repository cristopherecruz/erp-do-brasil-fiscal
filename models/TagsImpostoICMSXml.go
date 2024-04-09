package models

type TagsImpostoICMSXml struct {
	ICMS00    TagsImpostoICMS00Xml    `xml:"ICMS00,omitempty"`
	ICMS10    TagsImpostoICMS10Xml    `xml:"ICMS10,omitempty"`
	ICMS40    TagsImpostoICMS40Xml    `xml:"ICMS40,omitempty"`
	ICMS60    TagsImpostoICMS60Xml    `xml:"ICMS60,omitempty"`
	ICMSSN101 TagsImpostoICMSSN101Xml `xml:"ICMSSN101,omitempty"`
	ICMSSN102 TagsImpostoICMSSN102Xml `xml:"ICMSSN102,omitempty"`
	ICMSSN201 TagsImpostoICMSSN201Xml `xml:"ICMSSN201,omitempty"`
	ICMSSN202 TagsImpostoICMSSN202Xml `xml:"ICMSSN202,omitempty"`
	ICMSSN500 TagsImpostoICMSSN500Xml `xml:"ICMSSN500,omitempty"`
	ICMSSN900 TagsImpostoICMSSN900Xml `xml:"ICMSSN900,omitempty"`
}
