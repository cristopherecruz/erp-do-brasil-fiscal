package models

type TagVolumeTransportadorXml struct {
	QuantidadeVolume    int    `xml:"qVol,omitempty"`
	Especie             string `xml:"esp,omitempty"`
	PesoLiquido         string `xml:"pesoL,omitempty"`
	PesoBruto           string `xml:"pesoB,omitempty"`
	Marca               string `xml:"marca,omitempty"`
	NumeracaoDosVolumes string `xml:"nVol,omitempty"`
}
