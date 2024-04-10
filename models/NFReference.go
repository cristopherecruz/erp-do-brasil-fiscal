package models

import "encoding/xml"

type NFReference struct {
	XMLName      xml.Name       `xml:"Reference"`
	URI          string         `xml:"URI,omitempty,attr"`
	Transforms   []NFTransform  `xml:"Transforms>Transform,omitempty"`
	DigestMethod NFDigestMethod `xml:"DigestMethod,omitempty"`
	DigestValue  string         `xml:"DigestValue,omitempty"`
}
