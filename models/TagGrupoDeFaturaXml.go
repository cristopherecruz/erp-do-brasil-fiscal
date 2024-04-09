package models

type TagGrupoDeFaturaXml struct {
	NFat  string  `xml:"nFat,omitempty"`
	VOrig float64 `xml:"vOrig,omitempty"`
	VDesc float64 `xml:"vDesc,omitempty"`
	VLiq  float64 `xml:"vLiq,omitempty"`
}
