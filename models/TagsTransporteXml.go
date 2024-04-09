package models

type TagsTransporteXml struct {
	ModFrete   ModalidadeFrete           `xml:"modFrete"`
	Transporta TagTransportadorXml       `xml:"transporta,omitempty"`
	Vol        TagVolumeTransportadorXml `xml:"vol,omitempty"`
}
