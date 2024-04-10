package models

import "encoding/xml"

type NFCanonicalizationMethod struct {
	XMLName   xml.Name `xml:"CanonicalizationMethod"`
	Algorithm string   `xml:"Algorithm,omitempty,attr"`
}
