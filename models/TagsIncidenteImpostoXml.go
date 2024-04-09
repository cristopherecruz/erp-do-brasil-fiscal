package models

type TagsIncidenteImpostoXml struct {
	ICMS       TagsImpostoICMSXml       `xml:"ICMS,omitempty"`
	IPI        TagsImpostoIPIXml        `xml:"IPI,omitempty"`
	II         TagsImpostoImportacaoXml `xml:"II,omitempty"`
	PIS        TagsImpostoPISXml        `xml:"PIS,omitempty"`
	COFINS     TagsImpostoCOFINSXml     `xml:"COFINS,omitempty"`
	ICMSUFDest TagsICMSUFDestinoXml     `xml:"ICMSUFDest,omitempty"`
	VTotTrib   string                   `xml:"vTotTrib,omitempty"`
}
