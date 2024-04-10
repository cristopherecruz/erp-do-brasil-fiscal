package models

import "encoding/xml"

type NFSignature struct {
	XMLName        xml.Name     `xml:"Signature" namespace:"http://www.w3.org/2000/09/xmldsig#"`
	SignedInfo     NFSignedInfo `xml:"SignedInfo,omitempty"`
	SignatureValue string       `xml:"SignatureValue,omitempty"`
	KeyInfo        NFKeyInfo    `xml:"KeyInfo,omitempty"`
}
