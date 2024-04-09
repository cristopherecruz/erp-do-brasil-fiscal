package models

import "encoding/xml"

type TagsFiscalReferenciaXml struct {
	XMLName                         xml.Name                       `xml:"TagsFiscalReferenciaXml"`
	ChaveAcesso                     string                         `xml:"refNFe,omitempty"`
	Modelo1por1Referenciada         TagsModelo11AReferenciaXml     `xml:"refNF,omitempty"`
	InfoNFProdutorRuralReferenciada TagsProdutorRuralReferenciaXml `xml:"refNFP,omitempty"`
	ChaveAcessoCTReferenciada       string                         `xml:"refCTe,omitempty"`
	CupomFiscalReferenciado         TagsCupomFiscalReferenciaXml   `xml:"refECF,omitempty"`
}
