package models

type TagIPITribXml struct {
	CST  SituacaoTributariaPIS `xml:"CST,omitempty"`
	VBC  string                `xml:"vBC,omitempty"`
	PIPI string                `xml:"pIPI,omitempty"`
	VIPI string                `xml:"vIPI,omitempty"`
}
