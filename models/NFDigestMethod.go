package models

import "encoding/xml"

type NFDigestMethod struct {
	XMLName   xml.Name `xml:"DigestMethod"`
	Algorithm string   `xml:"Algorithm,omitempty,attr"`
}
