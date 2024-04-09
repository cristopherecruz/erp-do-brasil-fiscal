package models

import "encoding/xml"

type TagsProdutorRuralReferenciaXml struct {
	XMLName               xml.Name `xml:"TagsProdutorRuralReferenciaXml"`
	UFEmitente            string   `xml:"cUF,omitempty"`
	AnoMesEmissao         string   `xml:"AAMM,omitempty"`
	CNPJEmitente          string   `xml:"CNPJ,omitempty"`
	CPFEmitente           string   `xml:"CPF,omitempty"`
	IEEmitente            string   `xml:"IE,omitempty"`
	ModeloDocumentoFiscal string   `xml:"mod,omitempty"`
	SerieDocumentoFiscal  int      `xml:"serie,omitempty"`
	NumeroDocumentoFiscal int      `xml:"nNF,omitempty"`
}
