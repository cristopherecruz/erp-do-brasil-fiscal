package models

type TagsImpostoIPIXml struct {
	CEnq    string              `xml:"cEnq,omitempty"`
	IPITrib TagIPITribXml       `xml:"IPITrib,omitempty"`
	IPINT   TagsImpostoPISNTXml `xml:"IPINT,omitempty"`
}
