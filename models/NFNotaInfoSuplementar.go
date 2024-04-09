package models

import "encoding/xml"

type NFNotaInfoSuplementar struct {
	XMLName                xml.Name `xml:"infNFeSupl" namespace:"http://www.portalfiscal.inf.br/nfe"`
	QRCode                 string   `xml:",chardata"`
	URLConsultaChaveAcesso string   `xml:"urlChave"`
}
