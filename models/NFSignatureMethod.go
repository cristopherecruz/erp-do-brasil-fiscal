package models

import "encoding/xml"

type NFSignatureMethod struct {
	XMLName   xml.Name `xml:"SignatureMethod"`
	Algorithm string   `xml:"Algorithm,omitempty,attr"`
}
