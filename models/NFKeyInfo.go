package models

import "encoding/xml"

type NFKeyInfo struct {
	XMLName  xml.Name   `xml:"KeyInfo"`
	X509Data NFX509Data `xml:"X509Data,omitempty"`
}
