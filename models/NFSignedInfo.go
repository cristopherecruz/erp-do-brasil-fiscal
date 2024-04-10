package models

import "encoding/xml"

type NFSignedInfo struct {
	XMLName                xml.Name                 `xml:"SignedInfo"`
	CanonicalizationMethod NFCanonicalizationMethod `xml:"CanonicalizationMethod,omitempty"`
	SignatureMethod        NFSignatureMethod        `xml:"SignatureMethod,omitempty"`
	Reference              NFReference              `xml:"Reference,omitempty"`
}
