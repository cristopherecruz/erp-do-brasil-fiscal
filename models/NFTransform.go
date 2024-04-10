package models

import "encoding/xml"

type NFTransform struct {
	XMLName   xml.Name `xml:"Transform"`
	Algorithm string   `xml:"Algorithm,omitempty,attr"`
}
