package models

import "encoding/xml"

type TagsModelo11AReferenciaXml struct {
	XMLName               xml.Name `xml:"TagsModelo11AReferenciaXml"`
	UF                    string   `xml:"cUF,omitempty"`
	AnoMesEmissaoNFe      string   `xml:"AAMM,omitempty"`
	CNPJ                  string   `xml:"CNPJ,omitempty"`
	ModeloDocumentoFiscal string   `xml:"mod,omitempty"`
	Serie                 int      `xml:"serie,omitempty"`
	NumeroDocumentoFiscal string   `xml:"nNF,omitempty"`
}
