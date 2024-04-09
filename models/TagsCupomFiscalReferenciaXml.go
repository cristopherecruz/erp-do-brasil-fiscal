package models

import "encoding/xml"

type TagsCupomFiscalReferenciaXml struct {
	XMLName                     xml.Name `xml:"TagsCupomFiscalReferenciaXml"`
	ModeloDocumentoFiscal       string   `xml:"mod,omitempty"`
	NumeroOrdemSequencialECF    string   `xml:"nECF,omitempty"`
	NumeroContadorOrdemOperacao string   `xml:"nCOO,omitempty"`
}
