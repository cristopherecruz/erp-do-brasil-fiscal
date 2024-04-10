package models

import "encoding/xml"

type NFX509Data struct {
	XMLName         xml.Name `xml:"X509Data"`
	X509Certificate string   `xml:"X509Certificate,omitempty"`
}
