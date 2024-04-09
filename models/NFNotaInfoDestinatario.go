package models

type NFNotaInfoDestinatario struct {
	CNPJ                 string                    `xml:"CNPJ,omitempty"`
	CPF                  string                    `xml:"CPF,omitempty"`
	IdEstrangeiro        string                    `xml:"idEstrangeiro,omitempty"`
	XNome                string                    `xml:"xNome,omitempty"`
	XFant                string                    `xml:"xFant,omitempty"`
	EnderecoDestinatario TagsEnderecoEmitenteXml   `xml:"enderDest,omitempty"`
	IndIEDest            NFIndicadorIEDestinatario `xml:"indIEDest"`
	IE                   string                    `xml:"IE,omitempty"`
	ISUF                 string                    `xml:"ISUF,omitempty"`
	IM                   string                    `xml:"IM,omitempty"`
	Email                string                    `xml:"email,omitempty"`
}
